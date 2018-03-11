package telegramHelpers

import "strconv"

type MessageBuilder struct {
	Text string
}


func MessageBuilderInit() *MessageBuilder {
	return &MessageBuilder{Text : ""}
}


func (message *MessageBuilder) BoldText(text string) *MessageBuilder {
	message.Text += "<strong>" + text + "</strong>"
	return message
}


func (message *MessageBuilder) ItalicText(text string) *MessageBuilder {
	message.Text += "<i>" + text + "</i>"
	return message
}


func (message *MessageBuilder) CodeText(text string) *MessageBuilder {
	message.Text += "<code>" + text + "</code>"
	return message
}


func (message *MessageBuilder) NormalText(text string) *MessageBuilder {
	message.Text += text
	return message
}


func (message *MessageBuilder) MentionUser(id int64, name string) *MessageBuilder {
	message.Text += "<a href='tg://user?id=" + strconv.FormatInt(id, 10) + "'>" + name + "</a>"
	return message
}


func (message *MessageBuilder) NewRow() *MessageBuilder {
	message.Text += "\n"
	message.CodeText("-----------")
	message.Text += "\n"
	return message
}