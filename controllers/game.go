package controllers

import (
	"fmt"
	"mygame/choices"
	"mygame/models"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

var playerMap = make(map[string]*models.Player)

type Player struct {
	Territory int // 领土数
	Step      int // 当前步骤
}

type Choice struct {
	Story      string // 当前故事背景
	TextA      string // 选项 A
	TextB      string // 选项 B
	TerritoryA int    // 选择 A 后增加的领土
	TerritoryB int    // 选择 B 后的领土变化
}

var mut sync.Mutex

func getCurrentPlayer(token string) *models.Player {
	mut.Lock()
	defer mut.Unlock()
	if _, ok := playerMap[token]; !ok {
		playerMap[token] = &models.Player{
			Name:        token,
			Territory:   50,
			CurrentStep: 1,
			Result:      "",
		}
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
	player := getCurrentPlayer(token)               // 获取当前玩家状态
	choice := choices.GetChoice(player.CurrentStep) // 获取当前选项

	c.JSON(http.StatusOK, gin.H{
		"territory":    player.Territory,
		"current_step": player.CurrentStep,
		"story":        choice.Story,
		"choice_a":     choice.TextA,
		"choice_b":     choice.TextB,
		"token":        token,
	})

}

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
	currentChoice := choices.GetChoice(player.CurrentStep)
	if currentChoice.Story == "游戏结束" {
		c.JSON(http.StatusOK, gin.H{"message": "游戏已结束"})
		return
	}
	chText := ""
	if input.Choice == "A" {
		player.Territory += currentChoice.TerritoryA
		chText = currentChoice.TextA
	} else if input.Choice == "B" {
		player.Territory += currentChoice.TerritoryB
		chText = currentChoice.TextB
	} else {
		// 换你怎么做
	}

	player.CurrentStep++

	// 游戏结束判断
	if player.CurrentStep > 10 {
		if player.Territory >= 100 {
			player.Result = "胜利"
		} else {
			player.Result = "失败"
		}
	} else {
		player.Result = fmt.Sprintf("%s 此时你%s，因为你的行为，领土数变为：%d", currentChoice.Story, chText, player.Territory)
	}

	c.JSON(http.StatusOK, player)
}
