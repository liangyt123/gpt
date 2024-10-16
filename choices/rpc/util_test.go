package rpc

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {

	txt := `{
		"背景": "你是一位年轻的国王，刚继承王位，王国内部问题重重。边境战火未停，国内粮食短缺，百姓困苦。你的大臣们各有立场，时常提出相互矛盾的建议。你必须在100回合内赢得足够的爱戴值，巩固王位，否则王国将陷入动荡。",
		"当前爱戴值": 5,
		"剧情": "你站在王座上，大臣们聚集在大厅中，讨论着眼前的局势。一位资深大臣提出：“陛下，北方敌军蠢蠢欲动，防御是当前首要任务。”\\n另一位年轻的大臣则反对道：“陛下，民生为重，粮食供应不足，百姓怨声载道。我们必须先解决粮食危机。”\\n你站在决策的十字路口，必须做出抉择。你的每一步决定都将影响王国的未来，以及你在臣民中的地位。",
		"剧情图像": "[dalle request: \\\"An anime-style scene showing a young king sitting on a grand throne in a medieval castle. Surrounding him are advisors and officials debating fiercely. The king looks thoughtful, with one elderly advisor pointing towards a map while a younger official gestures towards a crowd of citizens outside. The atmosphere is tense, and the setting sun casts long shadows across the grand hall. Size: 1792x1024.\\\"]",
		"可选择的选项": [
			{
				"id": 1,
				"选项": "优先加强边境防御，减少敌军威胁。"
			},
			{
				"id": 2,
				"选项": "启动紧急粮食配给，暂时缓和民众的愤怒。"
			}
		]
	}`

	scene, err := ParseStory(txt)
	if err != nil {
		fmt.Println("解析出错:", err)
		return
	}

	fmt.Printf("scene %+v", scene)

}
