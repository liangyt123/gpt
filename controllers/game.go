package controllers

import (
	"fmt"
	"mygame/choices"
	"mygame/choices/rpc"
	"mygame/models"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

var playerMap = make(map[string]*models.Player)
var historyMap = make(map[string][]*rpc.HistoryChoice)

type Player struct {
	Territory int // 爱戴值
	Step      int // 当前步骤
}

var mut sync.Mutex

func getCurrentPlayer(token string) *models.Player {
	mut.Lock()
	defer mut.Unlock()
	if _, ok := playerMap[token]; !ok {

		player := &models.Player{
			Name:        token,
			Territory:   50,
			CurrentStep: 1,
			Result:      "",
			InitIntro:   "你是一个普通的人，你的爱戴值为50，你的选择将会影响你的爱戴值，当爱戴值大于100时，你胜利，当爱戴值小于0时，你失败",
		}
		playerMap[token] = player
		cli := rpc.Client{BaseURL: ""}
		r, err := cli.MockChoice(rpc.Req{
			Text:  "初始化",
			Story: "你是一个普通的人，你的爱戴值为50，你的选择将会影响你的爱戴值，当爱戴值大于100时，你胜利，当爱戴值小于0时，你失败",
			Round: 0,
			History: []*rpc.HistoryChoice{
				{
					Text:        "初始化",
					Round:       1,
					Territory:   50,
					Story:       "你是一个普通的人，你的爱戴值为50，你的选择将会影响你的爱戴值，当爱戴值大于100时，你胜利，当爱戴值小于0时，你失败",
					ImageBase64: "",
				},
			},
		})
		if err != nil {
			fmt.Println("Failed to get a valid response")
		}

		player.CurrentChoice = choices.Choice{
			TextA:       r.TextA,
			TextB:       r.TextB,
			TerritoryA:  r.TerritoryA,
			TerritoryB:  r.TerritoryB,
			Story:       r.Story,
			ImageBase64: r.ImageBase64,
		}
		fmt.Println("player", player)

	}
	return playerMap[token]
}

// func getCurrentChoiceList(token string) *choices.Choice {
// 	mut.Lock()
// 	defer mut.Unlock()
// 	if playerMap[token].CurrentStep-1 >= len(choiceMap[token]) {
// 		return &choices.Choice{
// 			TextA:      "游戏结束",
// 			TextB:      "游戏结束",
// 			TerritoryA: 0,
// 			TerritoryB: 0,
// 			Story:      "游戏结束",
// 		}
// 	}
// 	return &choiceMap[token][playerMap[token].CurrentStep-1]
// }

var random = rand.New(rand.NewSource(uint64(time.Now().Unix())))
var gToken = func() string {
	key := make([]byte, 10)
	for i := 0; i < 10; i++ {
		key[i] = byte(random.Intn(26) + 65)
	}
	return string(key)
}

// 获取玩家信息
func GetPlayerInfo(c *gin.Context) {
	var input struct {
		Token string `json:"token"`
	}
	c.BindJSON(&input)
	token := input.Token
	if token == "" {
		//生成 token
		token = gToken()

	}
	// 返回当前玩家信息和故事背景
	player := getCurrentPlayer(token) // 获取当前玩家状态
	choice := player.CurrentChoice

	c.JSON(http.StatusOK, gin.H{
		"territory":    player.Territory,
		"current_step": player.CurrentStep,
		"story":        choice.Story,
		"choice_a":     choice.TextA,
		"choice_b":     choice.TextB,
		"mini_game":    choice.MiniGame,
		"image_base64": player.ImageBase64,
		"token":        token,
	})

}

var falseEnd = "失败"
var trueEnd = "胜利"
var badEnd = "因为你的多次错误选择，爱戴值小于 0，你失败了，成为了一个🤡"
var goodEnd = "因为你的多次正确选择，爱戴值大于 100，你胜利了，成为了一个👑"

// 处理玩家的选择
func MakeChoice(c *gin.Context) {
	var input struct {
		Choice string `json:"choice"` // "A" 或 "B"
		Token  string `json:"token"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	player := getCurrentPlayer(input.Token)
	// 根据当前步骤和选择更新玩家信息
	currentChoice := player.CurrentChoice
	if player.Result == falseEnd || player.Result == trueEnd || player.Result == badEnd || player.Result == goodEnd || currentChoice.Story == "游戏结束" {
		c.JSON(http.StatusOK, gin.H{"message": "游戏已结束"})
		return
	}

	if player.Territory <= 0 {
		player.Result = badEnd
		c.JSON(http.StatusOK, player)
		return
	}
	if player.Territory >= 100 {
		player.Result = goodEnd
		c.JSON(http.StatusOK, player)
		return
	}

	chText := ""
	chVal := 0
	if input.Choice == "A" {
		player.Territory += currentChoice.TerritoryA
		chVal = currentChoice.TerritoryA
		chText = currentChoice.TextA
	} else if input.Choice == "B" {
		player.Territory += currentChoice.TerritoryB
		chVal = currentChoice.TerritoryB
		chText = currentChoice.TextB
	} else {
		// 换你怎么做
	}
	historyMap[input.Token] = append(historyMap[input.Token], &rpc.HistoryChoice{
		Text:        chText,
		Territory:   chVal,
		Story:       currentChoice.Story,
		ImageBase64: currentChoice.ImageBase64,
	})

	//发送请求获得新的选择

	player.CurrentStep++
	if player.Territory < 0 {
		player.Territory = 0
		player.Result = badEnd
		c.JSON(http.StatusOK, player)
		return
	}
	cli := rpc.Client{BaseURL: ""}
	r, err := cli.MockChoice(rpc.Req{
		Text:        chText,
		Story:       currentChoice.Story,
		ImageBase64: currentChoice.ImageBase64,
		Round:       player.CurrentStep,
		History:     historyMap[input.Token],
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get a valid response"})
		return
	}

	player.CurrentChoice = choices.Choice{
		TextA:       r.TextA,
		TextB:       r.TextB,
		TerritoryA:  r.TerritoryA,
		TerritoryB:  r.TerritoryB,
		Story:       r.Story,
		ImageBase64: r.ImageBase64,
	}

	// 游戏结束判断
	if player.CurrentStep > len(choices.Choices) {
		if player.Territory >= 100 {
			player.Result = trueEnd
		} else {
			player.Result = falseEnd
		}
	} else {
		player.Result = fmt.Sprintf("%s 此时你选择了%s，因为你的行为,爱戴值变为：%d", currentChoice.Story, chText, player.Territory)
	}

	c.JSON(http.StatusOK, player)
}
