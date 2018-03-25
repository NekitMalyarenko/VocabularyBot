package telegramServices

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"encoding/json"
	"log"
	"errors"
)


type ButtonData struct {
	Text     string
	funcId   string
	Data     map[string]interface{}
	IsNewRow bool
}


type KeyboardBuilder struct {
	hidden []ButtonData
}


func KeyboardBuilderInit() *KeyboardBuilder {
	return new(KeyboardBuilder)
}


func ButtonInit(button string) (data map[string]interface{},functionId string, err error) {
	err = json.Unmarshal([]byte(button), &data)
	if err != nil {
		log.Println(err)
		return nil,"", err
	}
	return data, data["funcId"].(string), nil
}


func (keyboard *KeyboardBuilder) NewButton(text string, isNewRow bool, functionId string) *KeyboardBuilder {
	button := ButtonData{
		Text     : text,
		IsNewRow : isNewRow,
		funcId   : functionId,
		Data     : make(map[string]interface{}, 0),
	}
	keyboard.hidden = append(keyboard.hidden, button)
	return keyboard
}


func (keyboard *KeyboardBuilder) PutData(data map[string]interface{}) *KeyboardBuilder {
	if data["funcId"] != nil {
		log.Println(errors.New("key funcId is already reserved"))
		return nil
	}

	keyboard.hidden[len(keyboard.hidden) - 1].Data = data
	return keyboard
}


func (button *ButtonData) getCallbackData() string {
	button.Data["funcId"] = button.funcId
	parsedData, err := json.Marshal(button.Data)
	if err != nil {
		log.Println(err)
		return ""
	}

	return string(parsedData)
}


func (keyboard *KeyboardBuilder) GetKeyboard() tgbotapi.InlineKeyboardMarkup {
	res := make([][]tgbotapi.InlineKeyboardButton, 0)

	for i := 0; i < len(keyboard.hidden); i++ {
		buttonData := keyboard.hidden[i]

		if buttonData.IsNewRow {
			res = append(res, tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(buttonData.Text, buttonData.getCallbackData())))
		} else {

			if len(res) != 0 {
				lastRowId := len(res) - 1
				res[lastRowId] = append(res[lastRowId], tgbotapi.NewInlineKeyboardButtonData(buttonData.Text, buttonData.getCallbackData()))
			} else {
				res = append(res, make([]tgbotapi.InlineKeyboardButton, 0))
				res[0] = append(res[0], tgbotapi.NewInlineKeyboardButtonData(buttonData.Text, buttonData.getCallbackData()))
			}
		}

	}

	return tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: res,
	}
}
