package main

import (
	"fmt"
	"math/rand"
)

func GenerateStoryEvent(p *Player) {
	story := fmt.Sprintf("在经历了数次战斗后，%s 决定向西扩张，他的 %s 率领一支军队进入新的领土...", p.Name, p.HandCards[0].Name)
	fmt.Println(story)
}

type Card struct {
	Name   string
	Power  int
	Skill  string
	Effect func(*Player) `json:"-"`
}

func GenerateCard(p *Player) Card {
	cardName := fmt.Sprintf("无双猛将 %s", p.HandCards[0].Name)
	cardPower := rand.Intn(20) + p.HandCards[0].Power
	cardEffect := func(p *Player) {
		fmt.Printf("%s 使用了卡牌 %s，获得了额外 %d 点战斗力！\n", p.Name, cardName, cardPower)
		p.HandCards[0].Power += cardPower
	}

	return Card{
		Name:   cardName,
		Power:  cardPower,
		Effect: cardEffect,
	}
}

func UseCard(p *Player) {
	card := GenerateCard(p)
	fmt.Printf("生成卡牌: %s (战斗力: %d)\n", card.Name, card.Power)
	card.Effect(p)
}
