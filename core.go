package main

import (
	"fmt"
	"math/rand"
)

// Player represents a game player
type Player struct {
	Name      string
	Resources int
	Army      int64
	HandCards []Card
}

// Event represents a game event
type Event struct {
	Description string
	Name        string `json:"name"`
	Effect      func(*Player)
	Type        string `json:"type"`
	Value       int    `json:"value,omitempty"`
}

// NewPlayer initializes a new player
func NewPlayer(name string) *Player {
	return &Player{
		Name:      name,
		Resources: 100,
		Army:      100,
		HandCards: []Card{},
	}
}

// RecruitGeneral adds a new general to the player's HandCards
func (p *Player) RecruitGeneral(g Card) {
	p.HandCards = append(p.HandCards, g)
}

// GenerateRandomEvent generates a random event
func GenerateRandomEvent() Event {
	eventOptions := []Event{
		{
			Description: "天气骤变，军队战斗力下降。",
			Effect:      func(p *Player) { p.Resources -= 10 }},
		{
			Description: "发现宝藏，资源大幅增加。",
			Effect:      func(p *Player) { p.Resources += 20 }},
	}
	return eventOptions[rand.Intn(len(eventOptions))]
}

// GenerateDynamicEvent generates a dynamic event based on the player's state
func GenerateDynamicEvent(p *Player) Event {
	if p.Resources < 30 {
		return Event{
			Description: "资源紧张，民众叛乱。",
			Effect: func(p *Player) {
				p.Resources -= 10
			},
		}
	} else if len(p.HandCards) > 2 {
		return Event{
			Description: "军队过大，内部出现分裂。",
			Effect: func(p *Player) {
				p.HandCards[0].Power -= 5
			},
		}
	} else {
		return GenerateRandomEvent()
	}
}

// HandleEvent applies an event to the player
func HandleEvent(p *Player) {
	event := GenerateDynamicEvent(p)
	fmt.Println(event.Description)
	event.Effect(p)
}

type Decision struct {
	Prompt  string
	Options map[string]func(*Player)
}

func MakeDecision(p *Player) {
	decision := Decision{
		Prompt: "敌军来袭，你有以下选择：",
		Options: map[string]func(*Player){
			"扩军": func(p *Player) {
				p.Resources -= 20
				newGeneral := drawCard(p, config.Cards)
				fmt.Printf("你选择了扩军，招募了武将 %s\n", newGeneral.Name)
			},
			"结盟": func(p *Player) {
				fmt.Printf("你选择了与附近势力结盟，获得了资源支援。\n")
				p.Resources += 20
			},
			"妥协": func(p *Player) {
				fmt.Printf("你选择了妥协，损失了部分资源，但避免了战争。\n")
				p.Resources -= 10
			},
		},
	}

	// 提示玩家选择
	fmt.Println(decision.Prompt)
	for option := range decision.Options {
		fmt.Printf("选择: %s\n", option)
	}

	// 模拟玩家输入，假设选择了扩军
	decision.Options["扩军"](p)
}

func GenerateDialogue(p *Player, g Card) {
	dialogue := fmt.Sprintf("%s 向 %s 询问了下一步的战略，%s 表示要进行突击！", p.Name, g.Name, g.Name)
	fmt.Println(dialogue)
}
