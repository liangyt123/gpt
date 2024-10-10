package models

import "mygame/choices"

// Player struct 定义主公信息
type Player struct {
	Name          string         `json:"name"`
	Territory     int            `json:"territory"`
	CurrentStep   int            `json:"current_step"`
	CurrentChoice choices.Choice `json:"current_choice"`
	Result        string         `json:"result"`
	ImageBase64   string         `json:"image_base64"`
	InitIntro     string         `json:"init_intro"`
}
