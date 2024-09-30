package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

var player *Player

// 配置结构体
type Config struct {
	Cards         []Card  `json:"cards"`
	Events        []Event `json:"events"`
	InitialPlayer struct {
		Name      string `json:"name"`
		Resources int    `json:"resources"`
		Army      int    `json:"army"`
	} `json:"initialPlayer"`
}

// 读取配置文件
func loadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// 随机事件
func randomEvent(player1, player2 *Player, events []Event) {
	event := events[rand.Intn(len(events))]
	fmt.Printf("发生了事件：%s - %s\n", event.Name, event.Description)

	switch event.Type {
	case "discard":
		// 玩家随机失去1-3张卡牌
		for _, player := range []*Player{player1, player2} {
			if len(player.HandCards) > 0 {
				removeCount := rand.Intn(3) + 1
				if removeCount > len(player.HandCards) {
					removeCount = len(player.HandCards)
				}
				player.HandCards = player.HandCards[:len(player.HandCards)-removeCount]
			}
		}
	case "boost":
		player1.Army += player1.Army / 5
		player2.Army += player2.Army / 5
		fmt.Printf("玩家1军队：%d, 玩家2军队：%d\n", player1.Army, player2.Army)
	case "steal":
		if player2.Resources >= event.Value {
			player1.Resources += event.Value
			player2.Resources -= event.Value
			fmt.Printf("%s 从 %s 窃取了 %d 金币\n", player1.Name, player2.Name, event.Value)
		}
	}
}

// 抽取卡牌
func drawCard(player *Player, cards []Card) Card {
	card := cards[rand.Intn(len(cards))]
	player.HandCards = append(player.HandCards, card)
	fmt.Printf("%s 抽取了卡牌：%s\n", player.Name, card.Name)
	return card
}

var config *Config

func main() {
	// 加载配置文件
	var err error
	config, err = loadConfig("config.json")
	if err != nil {
		fmt.Println("无法加载配置文件:", err)
		os.Exit(1)
	}

	player = NewPlayer("刘备")

	http.HandleFunc("/player", playerInfoHandler)
	http.HandleFunc("/draw-card", drawCardHandler)
	http.HandleFunc("/trigger-event", triggerEventHandler)
	http.HandleFunc("/use-card", userCardHandler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

// CORS middleware
func setupCORSResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*") // 允许所有来源
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func userCardHandler(w http.ResponseWriter, r *http.Request) {
	// 设置 CORS 头
	setupCORSResponse(&w, r)

	if r.Method == "OPTIONS" {
		return // 处理预检请求
	}

	w.Header().Set("Content-Type", "application/json")
	UseCard(player)
	json.NewEncoder(w).Encode(player.HandCards[0])
}

// Handler to get player information
func playerInfoHandler(w http.ResponseWriter, r *http.Request) {
	// 设置 CORS 头
	setupCORSResponse(&w, r)

	if r.Method == "OPTIONS" {
		return // 处理预检请求
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(player)
	json.NewEncoder(w).Encode(player)
}

// Handler to draw a card
func drawCardHandler(w http.ResponseWriter, r *http.Request) {
	// 设置 CORS 头
	setupCORSResponse(&w, r)

	if r.Method == "OPTIONS" {
		return // 处理预检请求
	}

	fmt.Println(config)
	g := drawCard(player, config.Cards)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(g)
	json.NewEncoder(w).Encode(
		struct {
			Name  string
			Power int
			Skill string
		}{
			Name:  g.Name,
			Power: g.Power,
			Skill: g.Skill,
		},
	)
}

// Handler to trigger a random event
func triggerEventHandler(w http.ResponseWriter, r *http.Request) {
	// 设置 CORS 头
	setupCORSResponse(&w, r)

	if r.Method == "OPTIONS" {
		return // 处理预检请求
	}
	player2 := NewPlayer("曹操")
	randomEvent(player, player2, config.Events)
	event := GenerateDynamicEvent(player)
	HandleEvent(player)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(event)
	type EventResponse struct {
		Description string `json:"Description"`
	}

	response := EventResponse{
		Description: event.Description,
	}

	json.NewEncoder(w).Encode(response)
}
