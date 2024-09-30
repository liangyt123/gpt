package choices

type Choice struct {
	TextA      string // 选项 A 文本
	TextB      string // 选项 B 文本
	TerritoryA int    // 选择 A 时领土变化
	TerritoryB int    // 选择 B 时领土变化
	Story      string // 当前选择的故事背景
}

var choices = []Choice{
	{
		TextA:      "寻求邻国的援助",
		TextB:      "组织内部的粮食分配",
		TerritoryA: 10,
		TerritoryB: -5,
		Story:      "公元 1200 年，饥荒席卷整个王国，粮食短缺导致人民不满。", // 场景 1
	},
	{
		TextA:      "派遣商人去邻国贸易",
		TextB:      "征收富人的财富来补贴穷人",
		TerritoryA: 5,
		TerritoryB: -10,
		Story:      "邻国同意提供帮助，但条件是分享贸易利益。", // 场景 2
	},
	{
		TextA:      "建造水库以收集雨水",
		TextB:      "加强对土地的耕种",
		TerritoryA: 15,
		TerritoryB: 5,
		Story:      "贸易使王国的经济有所回升，但水源匮乏。", // 场景 3
	},
	{
		TextA:      "招募农民一起耕作",
		TextB:      "引入外部技术提高产量",
		TerritoryA: 10,
		TerritoryB: 20,
		Story:      "新的水库解决了部分水源问题，但农民不足。", // 场景 4
	},
	{
		TextA:      "与教会合作进行宣传",
		TextB:      "鼓励人们团结一致，共度难关",
		TerritoryA: 8,
		TerritoryB: 12,
		Story:      "技术的引入大幅提高了产量，但社会动荡加剧。", // 场景 5
	},
	{
		TextA:      "引入新的作物",
		TextB:      "保持传统作物，保护文化",
		TerritoryA: 12,
		TerritoryB: 4,
		Story:      "宣传活动效果显著，但新的作物实验失败。", // 场景 6
	},
	{
		TextA:      "组织民众进行志愿活动",
		TextB:      "雇佣佣兵保护农业",
		TerritoryA: 10,
		TerritoryB: -5,
		Story:      "新的作物带来了短期的丰收，但志愿活动不足。", // 场景 7
	},
	{
		TextA:      "向王国的智者请教",
		TextB:      "寻求年轻人的建议",
		TerritoryA: 15,
		TerritoryB: 10,
		Story:      "农业保护措施加强，但经济持续低迷。", // 场景 8
	},
	{
		TextA:      "建立粮食储备系统",
		TextB:      "允许国际贸易以获取资源",
		TerritoryA: 20,
		TerritoryB: 10,
		Story:      "智者的建议使得王国开始建立粮食储备。", // 场景 9
	},
	{
		TextA:      "与地方贵族合作",
		TextB:      "进行改革，重新分配资源",
		TerritoryA: 10,
		TerritoryB: 15,
		Story:      "粮食储备系统稳固了王国的基础，但贵族的权力依然不可忽视。", // 场景 10
	},
}

// 根据步骤获取当前选择
func GetChoice(step int) Choice {
	if step > 0 && step <= len(choices) {
		return choices[step-1]
	}
	return Choice{Story: "游戏结束", TextA: "游戏结束", TextB: "游戏结束", TerritoryA: 0, TerritoryB: 0}
}
