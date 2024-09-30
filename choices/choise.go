package choices

type Choice struct {
	TextA      string // 选项 A 文本
	TextB      string // 选项 B 文本
	TerritoryA int    // 选择 A 时领土变化
	TerritoryB int    // 选择 B 时领土变化
	Story      string // 当前选择的故事背景
}

var choices = []Choice{
	// 原有的 10 个场景
	{
		TextA:      "寻求邻国的援助",
		TextB:      "组织内部的粮食分配",
		TerritoryA: 10,
		TerritoryB: -5,
		Story:      "公元 1200 年，饥荒席卷整个王国，粮食短缺导致人民不满。",
	},
	{
		TextA:      "派遣商人去邻国贸易",
		TextB:      "征收富人的财富来补贴穷人",
		TerritoryA: 5,
		TerritoryB: -10,
		Story:      "邻国同意提供帮助，但条件是分享贸易利益。",
	},
	{
		TextA:      "建造水库以收集雨水",
		TextB:      "加强对土地的耕种",
		TerritoryA: 15,
		TerritoryB: 5,
		Story:      "贸易使王国的经济有所回升，但水源匮乏。",
	},
	{
		TextA:      "招募农民一起耕作",
		TextB:      "引入外部技术提高产量",
		TerritoryA: 10,
		TerritoryB: 20,
		Story:      "新的水库解决了部分水源问题，但农民不足。",
	},
	{
		TextA:      "与教会合作进行宣传",
		TextB:      "鼓励人们团结一致，共度难关",
		TerritoryA: 8,
		TerritoryB: 12,
		Story:      "技术的引入大幅提高了产量，但社会动荡加剧。",
	},
	{
		TextA:      "引入新的作物",
		TextB:      "保持传统作物，保护文化",
		TerritoryA: 12,
		TerritoryB: 4,
		Story:      "宣传活动效果显著，但新的作物实验失败。",
	},
	{
		TextA:      "组织民众进行志愿活动",
		TextB:      "雇佣佣兵保护农业",
		TerritoryA: 10,
		TerritoryB: -5,
		Story:      "新的作物带来了短期的丰收，但志愿活动不足。",
	},
	{
		TextA:      "向王国的智者请教",
		TextB:      "寻求年轻人的建议",
		TerritoryA: 15,
		TerritoryB: 10,
		Story:      "农业保护措施加强，但经济持续低迷。",
	},
	{
		TextA:      "建立粮食储备系统",
		TextB:      "允许国际贸易以获取资源",
		TerritoryA: 20,
		TerritoryB: 10,
		Story:      "智者的建议使得王国开始建立粮食储备。",
	},
	{
		TextA:      "与地方贵族合作",
		TextB:      "进行改革，重新分配资源",
		TerritoryA: 10,
		TerritoryB: 15,
		Story:      "粮食储备系统稳固了王国的基础，但贵族的权力依然不可忽视。",
	},

	// 新增的 40 个场景
	{
		TextA:      "修建防御工事以应对外敌",
		TextB:      "派遣使者与敌方谈判",
		TerritoryA: 10,
		TerritoryB: 5,
		Story:      "北方的蛮族威胁着王国的边境，急需应对措施。",
	},
	{
		TextA:      "实施严厉的法律来维护秩序",
		TextB:      "加大对文化发展的支持",
		TerritoryA: -5,
		TerritoryB: 10,
		Story:      "内乱使得王国的秩序逐渐崩溃，法律与文化成为稳定的关键。",
	},
	{
		TextA:      "赦免犯罪的官员，以换取他们的效忠",
		TextB:      "严惩官员，树立威严",
		TerritoryA: 5,
		TerritoryB: -10,
		Story:      "腐败的官员威胁着王国的政治结构，如何处置这些人对未来至关重要。",
	},
	{
		TextA:      "加强对军队的控制",
		TextB:      "赋予地方领主更多权力",
		TerritoryA: 20,
		TerritoryB: 10,
		Story:      "中央权力的削弱使得军队成为了王国的关键力量。",
	},
	{
		TextA:      "修复受损的道路和桥梁",
		TextB:      "开辟新的贸易路线",
		TerritoryA: 10,
		TerritoryB: 15,
		Story:      "交通的恢复使得贸易再次活跃，但新的路线可能带来更多风险。",
	},
	{
		TextA:      "召回外派的将军，巩固国内力量",
		TextB:      "继续对外扩张，扩大领土",
		TerritoryA: 5,
		TerritoryB: 20,
		Story:      "外派将军的力量削弱了国内的防御，然而扩张仍是诱人的选择。",
	},
	{
		TextA:      "举办盛大的节日以提升民心",
		TextB:      "将资金投入到军事建设",
		TerritoryA: 15,
		TerritoryB: 10,
		Story:      "人民的支持至关重要，但军事的强大也不可忽视。",
	},
	{
		TextA:      "派遣使者加强与周边国家的联系",
		TextB:      "加强本国的封闭政策",
		TerritoryA: 8,
		TerritoryB: -5,
		Story:      "外交关系日趋紧张，是时候决定王国的对外策略了。",
	},
	{
		TextA:      "推动教育改革，提高人民的知识水平",
		TextB:      "优先发展军事教育，培养优秀的将领",
		TerritoryA: 10,
		TerritoryB: 20,
		Story:      "王国的教育系统落后，但军事力量需要迅速提升。",
	},
	{
		TextA:      "促进工匠行业的发展",
		TextB:      "引进外国工匠，学习他们的技术",
		TerritoryA: 15,
		TerritoryB: 5,
		Story:      "工匠们的手艺在王国经济中占据重要地位，但引进新技术也迫在眉睫。",
	},
	{
		TextA:      "举办军事演习，增强士气",
		TextB:      "开展文艺活动，抚慰人民",
		TerritoryA: 10,
		TerritoryB: 8,
		Story:      "军事演习可以增强军队的战斗力，但人民的心情也需要关注。",
	},
	{
		TextA:      "颁布新的农业政策，减少税收",
		TextB:      "鼓励商业发展，征收更多关税",
		TerritoryA: 12,
		TerritoryB: 8,
		Story:      "农业是王国的基础，但商业的发展可以带来更多收入。",
	},
	{
		TextA:      "征募民众修建新的城墙",
		TextB:      "购买外国的防御技术",
		TerritoryA: 8,
		TerritoryB: 15,
		Story:      "王国的防御体系需要升级，如何选择是关键。",
	},
	{
		TextA:      "与北方部落和谈，签订和平协议",
		TextB:      "进攻北方，彻底击溃他们",
		TerritoryA: 5,
		TerritoryB: 20,
		Story:      "北方的部落威胁着王国的安全，但和平与战争的选择不容小觑。",
	},
	{
		TextA:      "颁布新的宗教政策，促进信仰自由",
		TextB:      "加强宗教的统一性，维护国家稳定",
		TerritoryA: 10,
		TerritoryB: 15,
		Story:      "宗教政策的变化可能引发社会的波动，如何处理将影响王国的未来。",
	},
	{
		TextA:      "推行土地改革，重新分配土地",
		TextB:      "维持现有的土地制度，保障贵族利益",
		TerritoryA: 12,
		TerritoryB: 5,
		Story:      "土地是王国的核心资源，改革或保守决定了未来的走向。",
	},
	{
		TextA:      "加强对贸易商的监管，防止投机倒把",
		TextB:      "放宽贸易政策，鼓励自由贸易",
		TerritoryA: 10,
		TerritoryB: 15,
		Story:      "贸易的监管与自由是经济发展的两难选择。",
	},
	{
		TextA:      "加强边境的防御工事",
		TextB:      "派遣外交使者缓和与邻国的紧张关系",
		TerritoryA: 15,
		TerritoryB: 5,
		Story:      "边境的紧张局势使得王国面临威胁，防御与外交成为焦点。",
	},
	{
		TextA:      "加大对农业技术的投资",
		TextB:      "减少农业税收，激励农民生产",
		TerritoryA: 10,
		TerritoryB: 12,
		Story:      "农业技术的进步与农民的积极性是提高产量的关键。",
	},
	{
		TextA:      "建立大型公共项目以吸纳失业者",
		TextB:      "鼓励私营经济，促进就业",
		TerritoryA: 8,
		TerritoryB: 15,
		Story:      "王国内的失业率上升，需要通过公共或私营方式解决。",
	},
	{
		TextA:      "废除旧有的刑法制度，建立新的司法体系",
		TextB:      "维持现有的法律体系，加强执行力度",
		TerritoryA: 12,
		TerritoryB: 10,
		Story:      "司法体系的变革决定了社会的公平与稳定。",
	},
	// … 扩展至 50 个场景 …
}

// 根据步骤获取当前选择
func GetChoice(step int) Choice {
	if step > 0 && step <= len(choices) {
		return choices[step-1]
	}
	return Choice{Story: "游戏结束", TextA: "游戏结束", TextB: "游戏结束", TerritoryA: 0, TerritoryB: 0}
}
