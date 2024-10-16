package rpc

import (
	"bytes"
	"fmt"

	"encoding/json"
	"errors"
	"net/http"
)

// 实现 APIClient 接口
type Client struct {
}

type HistoryChoice []string

// 定义请求结构体

type Req struct {
	Text        string          `json:"text"`         // 选项 文本
	Story       string          `json:"story"`        // 选择的故事
	ImageBase64 string          `json:"image_base64"` // 此时生成的图片
	Round       int             `json:"round"`        // 当前回合数
	History     []HistoryChoice `json:"history"`      // 历史选择
}

// 定义响应结构体
type Resp struct {
	TextA      string `json:"text_a"`     // 选项 A 文本
	TextB      string `json:"text_b"`     // 选项 B 文本
	Text       string `json:"text"`       // 文本
	Background string `json:"background"` // 背景
	Territory  int    `json:"territory"`  // 当前的领地值
	Story      string `json:"story"`      // 选择后的故事
	ImageURL   string `json:"image_url"`  // 选择后生成的图片
	MiniGame   string `json:"mini_game"`  // 小游戏
}

var baseURL = "http://127.0.0.1:8000"

func (c *Client) MockChoice(req Req) (Resp, error) {
	// // 模拟返回数据
	// if time.Now().Unix()%10 == 1 {
	// 	// 这个是游戏事件
	// 	return Resp{
	// 		TextA:    "成功通过游戏，获得奖励",
	// 		TextB:    "不玩游戏",
	// 		Story:    "突然遇到神仙，看你气宇非凡，万中无一的君主，规定时间通过连连看有奖励，能获得奖励",
	// 		MiniGame: "连连看",
	// 	}, nil
	// }

	var o Story
	fmt.Printf("req %+v", req)
	var i = 0
	for i = 0; i < 3; i++ {
		resp, err := c.MakeChoice(req)
		if err != nil {
			fmt.Println("MakeChoice err", err)
		}
		fmt.Println("resp", resp)
		o, err = ParseStory(resp.Text)
		if err != nil {
			fmt.Println("ParseScene err", err)
		}
		fmt.Println("tmp", o)
		if checkStoryHasVal(o) {
			break
		}
	}
	if i == 3 {
		return Resp{
			TextA:      "",
			TextB:      "",
			Background: "",
			Territory:  0,
			Story:      "",
		}, errors.New("no valid scene")
	}

	return Resp{
		TextA:      o.Options[0].Choice,
		TextB:      o.Options[1].Choice,
		Background: o.Background,
		Territory:  o.CurrentAdoration,
		Story:      o.Plot,
	}, nil

}

func (c *Client) MakeChoice(req Req) (Resp, error) {
	url := baseURL + "/chat/"

	// 序列化请求数据
	jsonData, err := json.Marshal(req)
	if err != nil {
		return Resp{}, err
	}
	fmt.Println("jsonData", string(jsonData))

	// 发送 POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return Resp{}, err
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return Resp{}, errors.New("failed to get a valid response")
	}

	// 读取响应
	var respData Resp
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return Resp{}, err
	}

	return respData, nil
}
