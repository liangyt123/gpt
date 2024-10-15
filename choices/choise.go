package choices

type Choice struct {
	TextA      string `json:"text_a"`      // 选项 A 文本
	TextB      string `json:"text_b"`      // 选项 B 文本
	TerritoryA int    `json:"territory_a"` // 选择 A 时爱戴变化
	TerritoryB int    `json:"territory_b"` // 选择 B 时爱戴变化
	Story      string `json:"story"`       // 选择的故事
	Text       string `json:"text"`        // 选项文本
	MiniGame   string `json:"mini_game"`   // 小游戏
	ImageURL   string `json:"image_url"`   // 图片 URL
}

var EasyChoices = []Choice{
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
		TextB:      "继续对外扩张，扩大爱戴",
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
}

var Choices = []Choice{

	{
		TextA:      "寻求邻国的援助，但可能失去部分主权",
		TextB:      "组织内部的粮食分配，导致不满的民众",
		TerritoryA: 5,
		TerritoryB: -10,
		Story:      "公元 1200 年，饥荒席卷整个王国，人民不满情绪加剧。", // 场景 1
	},
	{
		TextA:      "派遣商人去邻国贸易，可能遭遇劫匪",
		TextB:      "征收富人的财富来补贴穷人，导致富人反抗",
		TerritoryA: 0,
		TerritoryB: -15,
		Story:      "邻国的贸易条件苛刻，内部矛盾加剧。", // 场景 2
	},
	{
		TextA:      "建造水库以收集雨水，面临财政危机",
		TextB:      "加强对土地的耕种，但可能导致土壤退化",
		TerritoryA: 10,
		TerritoryB: -5,
		Story:      "水源问题解决，财政却陷入困境。", // 场景 3
	},
	{
		TextA:      "招募农民一起耕作，但可能导致不平等",
		TextB:      "引入外部技术，可能造成文化冲突",
		TerritoryA: 5,
		TerritoryB: -10,
		Story:      "外部技术的引入带来经济改善，却引发不满。", // 场景 4
	},
	{
		TextA:      "与教会合作进行宣传，导致宗教冲突",
		TextB:      "鼓励人们团结一致，但可能出现极端分子",
		TerritoryA: 10,
		TerritoryB: -5,
		Story:      "宣传活动的推行，面临着信仰的挑战。", // 场景 5
	},
	{
		TextA:      "引入新的作物，可能引发粮食危机",
		TextB:      "保持传统作物，但可能导致产量下降",
		TerritoryA: -5,
		TerritoryB: 10,
		Story:      "新作物实验失败，民众恐慌。", // 场景 6
	},
	{
		TextA:      "组织民众进行志愿活动，但面临安全隐患",
		TextB:      "雇佣佣兵保护农业，导致财政负担加重",
		TerritoryA: 0,
		TerritoryB: -10,
		Story:      "志愿活动效果有限，佣兵费用昂贵。", // 场景 7
	},
	{
		TextA:      "向王国的智者请教，但可能得不到实用建议",
		TextB:      "寻求年轻人的建议，可能缺乏经验",
		TerritoryA: 0,
		TerritoryB: -5,
		Story:      "智者的智慧有限，年轻人的建议不成熟。", // 场景 8
	},
	{
		TextA:      "建立粮食储备系统，可能导致腐败",
		TextB:      "允许国际贸易以获取资源，但面临竞争压力",
		TerritoryA: 10,
		TerritoryB: -5,
		Story:      "储备系统建立，潜藏着管理上的风险。", // 场景 9
	},
	{
		TextA:      "与地方贵族合作，可能导致权力失衡",
		TextB:      "进行改革，可能遭遇贵族反抗",
		TerritoryA: 5,
		TerritoryB: -15,
		Story:      "贵族的利益与改革的目标，存在巨大矛盾。", // 场景 10
	},
	{
		TextA:      "鼓励学者进行技术创新，投资风险高",
		TextB:      "引进外国技术，但可能影响本国的技术发展",
		TerritoryA: 0,
		TerritoryB: -10,
		Story:      "技术的创新带来潜在的利益，但风险也巨大。", // 场景 11
	},
	{
		TextA:      "与地方势力联合，或引发更大的冲突",
		TextB:      "进行改革，可能引起地方势力的反抗",
		TerritoryA: 10,
		TerritoryB: -15,
		Story:      "合作的安全与改革的风险，如何选择？", // 场景 12
	},
	{
		TextA:      "开发新的农业技术，可能遭遇自然灾害",
		TextB:      "引进外部技术，可能引起文化冲突",
		TerritoryA: 0,
		TerritoryB: -5,
		Story:      "外部技术的引进带来好处，但也伴随风险。", // 场景 13
	},
	{
		TextA:      "与商人合作，可能导致国家控制力减弱",
		TextB:      "鼓励地方自主经营，可能导致经济分裂",
		TerritoryA: 5,
		TerritoryB: -10,
		Story:      "经济合作与控制，哪一种更能提升国家的实力？", // 场景 14
	},
	{
		TextA:      "加强对农民的培训，但可能产生反效果",
		TextB:      "发展大型农场，导致农民失业",
		TerritoryA: 0,
		TerritoryB: -15,
		Story:      "培训与发展的矛盾，农民的命运如何？", // 场景 15
	},
	{
		TextA:      "借鉴魏国的治国理念，可能导致民众反感",
		TextB:      "学习蜀汉的团结与互助，但可能失去灵活性",
		TerritoryA: 0,
		TerritoryB: -5,
		Story:      "不同国家的经验，如何在现实中执行？", // 场景 16
	},
	{
		TextA:      "与百姓建立信任关系，但可能遭遇背叛",
		TextB:      "加强对百姓的管理，可能引发不满",
		TerritoryA: 10,
		TerritoryB: -5,
		Story:      "信任与控制之间的平衡，王国应如何选择？", // 场景 17
	},
	{
		TextA:      "建立庙会，可能引发宗教冲突",
		TextB:      "推进官府文化教育，但可能引起反感",
		TerritoryA: 0,
		TerritoryB: -10,
		Story:      "文化活动与政府教育的冲突，如何选择以平衡？", // 场景 18
	},
	{
		TextA:      "引导年轻人参加农耕活动，但可能影响他们的未来",
		TextB:      "鼓励年轻人参与军事训练，可能导致和平破裂",
		TerritoryA: 5,
		TerritoryB: -10,
		Story:      "年轻人未来的方向与王国的安全，如何抉择？", // 场景 19
	},
	{
		TextA:      "投资水利工程，可能导致资金短缺",
		TextB:      "建设储粮仓库，但可能引发腐败问题",
		TerritoryA: 5,
		TerritoryB: -15,
		Story:      "基础设施的建设伴随的经济负担，如何选择？", // 场景 20
	},
	{
		TextA:      "派遣特使与各大势力谈判，可能遭遇背叛",
		TextB:      "组织大规模的军演，可能引起邻国的不满",
		TerritoryA: 10,
		TerritoryB: -5,
		Story:      "外交与军事的对抗，王国应如何选择？", // 场景 21
	},
	{
		TextA:      "建立相互合作的商会，可能导致权力失衡",
		TextB:      "实施国家控制的经济政策，可能引发民众不满",
		TerritoryA: 5,
		TerritoryB: -10,
		Story:      "商会与国家控制的矛盾，如何选择以推动经济？", // 场景 22
	},
	{
		TextA:      "与各地英雄豪杰交流，但可能遭遇权力斗争",
		TextB:      "培养自己的将领，可能导致对外敌人的误解",
		TerritoryA: 10,
		TerritoryB: -5,
		Story:      "英雄豪杰的聚集与培养，王国如何应对内外压力？", // 场景 23
	},
	{
		TextA:      "利用间谍获取情报，但可能引发信任危机",
		TextB:      "通过公开透明的方式建立关系，可能失去机会",
		TerritoryA: 0,
		TerritoryB: -10,
		Story:      "情报的获取与信任的建立，如何选择以获取优势？", // 场景 24
	},

	{
		TextA: "成功通过游戏，获得奖励",
		TextB: "不玩游戏",

		TerritoryA: 10,
		TerritoryB: 0,
		Story:      "突然遇到神仙，看你气宇非凡，万中无一的君主，规定时间通过连连看有奖励，能获得奖励", // 场景 25
		MiniGame:   "连连看",
	},
}

// // 根据步骤获取当前选择
// func GetChoice(step int) Choice {
// 	if step > 0 && step <= len(Choices) {
// 		return Choices[step-1]
// 	}
// 	return Choice{Story: "游戏结束", TextA: "游戏结束", TextB: "游戏结束", TerritoryA: 0, TerritoryB: 0}
// }
