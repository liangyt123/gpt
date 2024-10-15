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
	BaseURL string
}

// 定义请求结构体
type HistoryChoice struct {
	Text        string `json:"text"`         // 选项 文本
	Territory   int    `json:"territory"`    // 选择时爱戴变化
	Story       string `json:"story"`        // 选择的故事
	ImageBase64 string `json:"image_base64"` // 此时生成的图片
	Round       int    `json:"round"`        // 当前回合数
}

type Req struct {
	Text        string           `json:"text"`         // 选项 文本
	Story       string           `json:"story"`        // 选择的故事
	ImageBase64 string           `json:"image_base64"` // 此时生成的图片
	Round       int              `json:"round"`        // 当前回合数
	History     []*HistoryChoice `json:"history"`      // 历史选择
}

// 定义响应结构体
type Resp struct {
	TextA       string `json:"text_a"`       // 选项 A 文本
	TextB       string `json:"text_b"`       // 选项 B 文本
	TerritoryA  int    `json:"territory_a"`  // 选择 A 时爱戴变化
	TerritoryB  int    `json:"territory_b"`  // 选择 B 时爱戴变化
	Story       string `json:"story"`        // 选择的故事
	ImageBase64 string `json:"image_base64"` // 生成图片
}

func (c *Client) MockChoice(req Req) (Resp, error) {
	// 模拟返回数据

	return c.MakeChoice(req)
}

func (c *Client) MakeChoice(req Req) (Resp, error) {
	c.BaseURL = "http://127.0.0.1:8000"
	url := c.BaseURL + "/chat/"

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
