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
var historyMap = make(map[string][]rpc.HistoryChoice)

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
			Territory:   5,
			CurrentStep: 0,
			Result:      "",
		}
		playerMap[token] = player
		cli := rpc.Client{}
		req := rpc.Req{
			Round: 0,
		}
		req.History = []rpc.HistoryChoice{}
		r, err := cli.MakeChoice2(req)
		if err != nil {
			fmt.Println("Failed to get a valid response")
		}
		fmt.Printf("%+v", r)

		player.CurrentChoice = &choices.Choice{
			TextA:      r.TextA,
			TextB:      r.TextB,
			ResultA:    r.ResultA,
			ResultB:    r.ResultB,
			Text:       r.Text,
			Territory:  r.Territory,
			Story:      r.Story,
			Background: r.Background,
		}
		historyMap[token] = []rpc.HistoryChoice{
			{"开始游戏", r.Text},
		}
		fmt.Printf("player +%v \n", player)

	}
	return playerMap[token]
}

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
		"token":        token,
		"background":   choice.Background,
	})

}

var endStep = 30
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
	if player.Territory >= endStep {
		player.Result = goodEnd
		c.JSON(http.StatusOK, player)
		return
	}

	chText := ""
	chNo := 0
	chResult := ""
	thisText := player.CurrentChoice.Text

	if input.Choice == "A" {
		chText = currentChoice.TextA
		chResult = currentChoice.ResultA
		chNo = 1
	} else if input.Choice == "B" {
		chText = currentChoice.TextB
		chResult = currentChoice.ResultB
		chNo = 2
	} else {
		// 换你怎么做
	}

	//发送请求获得新的选择

	player.CurrentStep++
	if player.Territory < 0 {
		player.Territory = 0
		player.Result = badEnd
		c.JSON(http.StatusOK, player)
		return
	}
	cli := rpc.Client{}

	r, err := cli.MakeChoice2(rpc.Req{
		Text:    chText,
		Story:   currentChoice.Story,
		Round:   player.CurrentStep,
		History: historyMap[input.Token],
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get a valid response"})
		return
	}

	historyMap[input.Token] = append(historyMap[input.Token], rpc.HistoryChoice{
		fmt.Sprintf("%d", chNo),
		thisText,
	})
	player.Territory = r.Territory
	player.CurrentChoice = &choices.Choice{
		TextA:      r.TextA,
		TextB:      r.TextB,
		ResultA:    r.ResultA,
		ResultB:    r.ResultB,
		Text:       r.Text,
		Story:      r.Story,
		Background: r.Background,
		Territory:  r.Territory,
		MiniGame:   r.MiniGame,
	}

	// 游戏结束判断
	if player.CurrentStep > endStep {
		if player.Territory >= 100 {
			player.Result = trueEnd
		} else {
			player.Result = falseEnd
		}
	} else {
		player.Result = fmt.Sprintf("%s 此时你选择了%s 结果:%s 因为你的行为，爱戴值变为：%d", currentChoice.Story, chText, chResult, player.Territory)
	}

	c.JSON(http.StatusOK, player)
}

// 处理 POST 请求的函数
func GenerateImage(c *gin.Context) {

	// ReqBody 定义请求体结构
	type ReqBody struct {
		Token string `json:"token"` // token 字段为必填
		Story string `json:"story"` // story 字段为必填
	}
	var reqBody ReqBody

	// 绑定 JSON 请求体到结构体
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 模拟生成返回的数据
	player := getCurrentPlayer(reqBody.Token)
	r1, _ := rpc.GenImage(reqBody.Story, rand.Intn(10)%3)
	if r1.ImgURL != "" {
		mut.Lock()
		player.CurrentChoice.ImgBase64 = r1.ImgBase64
		player.CurrentChoice.ImgPrompt = r1.ImgPrompt
		player.CurrentChoice.ImgURL = r1.ImgURL

		mut.Unlock()
	}
	fmt.Println("player", player)

	// 返回 JSON 响应
	c.JSON(http.StatusOK, player)
}
