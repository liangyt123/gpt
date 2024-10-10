



请求结构体

type Choice struct {

    Text       string`json:"text"`      // 选项 文本

    Territory  int`json:"territory"` // 选择时爱戴变化

    Story       string`json:"story"`       // 选择的故事

    ImageBase64 string`json:"image_base64"` 此时生成的图片

}

type req struct {

    Text       string`json:"text"`      // 选项 文本

     Story       string `json:"story"`
    // 选择的故事

    ImageBase64 string`json:"image_base64"` 此时生成的图片

    Round int`json:"round"`

    History []Choice`json:"history"`

}

返回结构体

type resp struct {

    TextA       string`json:"text_a"`      // 选项 A 文本

    TextB       string`json:"text_b"`      // 选项 B 文本

    TerritoryA  int`json:"territory_a"` // 选择 A 时爱戴变化

    TerritoryB  int`json:"territory_b"` // 选择 B 时爱戴变化

    Story       string`json:"story"`       // 选择的故事

    ImageBase64 string`json:"image_base64"` 生成图片

}



你可以使用 `curl` 或 Postman 测试此 API，发送类似以下的请求：

`curl -X POST http://localhost:8080/api/choice \

  -H "Content-Type: application/json" \

  -d '{

    "text": "测试文本",

    "story": "测试故事",

    "image_base64": "base64图片字符串",

    "round": 1,

    "history": [

    {

    "text": "历史选项1",

    "territory": 5,

    "story": "历史故事1",

    "image_base64": "base64图片字符串1"

    }

    ]

    }'`

你将收到类似这样的响应：

`{     "text_a": "选项 A 文本",     "text_b": "选项 B 文本",     "territory_a": 10,     "territory_b": -5,     "story": "测试故事",     "image_base64": "base64图片字符串" }`
