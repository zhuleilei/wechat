package message

//Text 文本消息
type Text struct {
	CommonToken
	Content string `xml:"Content" json:"-"`
	Text struct{
		Content string `json:"content" xml:"content"`
	} `json:"text" xml:"text"`
}

//NewText 初始化文本消息
func NewText(content string) *Text {
	text := new(Text)
	text.Text.Content = content
	return text
}
