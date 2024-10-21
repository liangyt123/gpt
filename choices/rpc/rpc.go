package rpc

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"encoding/json"
	"errors"
	"net/http"
)

var HOST = "http://192.168.116.35"

// 实现 APIClient 接口
type Client struct {
}

type HistoryChoice []string

// 定义请求结构体

type Req struct {
	Text                string          `json:"text"`                  // 选项 文本
	Story               string          `json:"story"`                 // 选择的故事
	ImageBase64         string          `json:"image_base64"`          // 此时生成的图片
	Round               int             `json:"round"`                 // 当前回合数
	History             []HistoryChoice `json:"history"`               // 历史选择
	ShouldGenerateImage bool            `json:"should_generate_image"` // 是否生成图片
	Server              int             `json:"server"`                // 服务器
}

// 定义响应结构体
type Resp struct {
	TextA      string `json:"text_a"`     // 选项 A 文本
	TextB      string `json:"text_b"`     // 选项 B 文本
	ResultA    string `json:"result_a"`   // 选项 A 结果
	ResultB    string `json:"result_b"`   // 选项 B 结果
	Text       string `json:"text"`       // 文本
	Background string `json:"background"` // 背景
	Territory  int    `json:"territory"`  // 当前的领地值
	Story      string `json:"story"`      // 选择后的故事
	MiniGame   string `json:"mini_game"`  // 小游戏

}

var baseURL = HOST + ":8000"
var baseImageURL = HOST + ":8001"

func (c *Client) MakeChoice2(req Req) (Resp, error) {
	// // 模拟返回数据

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
		ResultA:    o.Options[0].Result,
		ResultB:    o.Options[1].Result,
		Background: o.Background,
		Text:       o.Text,
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

type ReqImageBody struct {
	Story     string `json:"story"`      // 请求的故事内容
	ImgServer int    `json:"img_server"` // 图片服务器ID
}

type ReturnImageBody struct {
	ImgPrompt string `json:"img_prompt"`           // 图片提示
	ImgBase64 string `json:"img_base64,omitempty"` // 图片的Base64编码，omitempty 表示如果为空则不返回
	ImgURL    string `json:"img_url,omitempty"`    // 图片的URL，omitempty 表示如果为空则不返回
}

func GenImage(story string, imgServer int) (ReturnImageBody, error) {
	// 构造请求体数据
	reqBody := ReqImageBody{
		Story:     story,
		ImgServer: imgServer,
	}
	// 将请求体序列化为 JSON
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error marshaling request body:", err)
		return ReturnImageBody{}, err
	}

	// 创建 HTTP POST 请求
	url := baseImageURL + "/gen_img/" // 替换为实际的 API URL
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ReturnImageBody{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ReturnImageBody{}, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ReturnImageBody{}, err
	}

	// 检查 HTTP 响应状态
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status code: %d\n", resp.StatusCode)
		fmt.Println("Response body:", string(body))
		return ReturnImageBody{}, errors.New("request failed")
	}

	// 解析响应体
	var returnBody ReturnImageBody
	err = json.Unmarshal(body, &returnBody)
	if err != nil {
		fmt.Println("Error unmarshaling response body:", err)
		return ReturnImageBody{}, err
	}
	fmt.Printf("Response body: %+v\n", returnBody)
	return returnBody, nil

}
