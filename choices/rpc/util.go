package rpc

import (
	"encoding/json"
	"fmt"
)

// 定义结构体
type Option struct {
	ID     int    `json:"id"`
	Choice string `json:"选项"`
	Result string `json:"结果"`
}

type Story struct {
	Background       string   `json:"背景"`
	CurrentAdoration int      `json:"当前爱戴值"`
	Plot             string   `json:"剧情"`
	PlotImage        string   `json:"剧情图像"`
	Options          []Option `json:"可选择的选项"`
	Text             string

	ImgPrompt  string `json:"img_prompt"`  // 图片提示
	ImgBase64  string `json:"img_base64"`  // 图片 base64
	ImgURL     string `json:"img_url"`     // 图片 url
	ImgStatues string `json:"img_statues"` // 图片状态
}

func ParseStory(data string) (Story, error) {
	// 创建 Story 结构体实例
	var story Story
	story.Text = data

	// 解析 JSON 数据
	err := json.Unmarshal([]byte(data), &story)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return story, err
	}
	return story, nil

}

func checkStoryHasVal(scene Story) bool {
	if scene.Background == "" || scene.Plot == "" || scene.CurrentAdoration == 0 || len(scene.Options) < 2 {
		return false
	}
	return true
}
