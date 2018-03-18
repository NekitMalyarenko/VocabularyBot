package telegramHelpers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)


type ButtonData struct {
	Text     string
	FuncId   int
	IsNewRow bool
}


type KeyboardBuilder struct {
	hidden []ButtonData
}


func KeyboardBuilderInit() *KeyboardBuilder {
	return new(KeyboardBuilder)
}


func (keyboard *KeyboardBuilder) NewButton(text string, isNewRow bool, funcId int) *KeyboardBuilder {
	button := ButtonData{
		Text:     text,
		IsNewRow: isNewRow,
		FuncId:   funcId,
	}
	keyboard.hidden = append(keyboard.hidden, button)
	return keyboard
}


func (keyboard *KeyboardBuilder) GetKeyboard() tgbotapi.InlineKeyboardMarkup {
	res := make([][]tgbotapi.InlineKeyboardButton, 0)

	for i := 0; i < len(keyboard.hidden); i++ {
		buttonData := keyboard.hidden[i]

		if buttonData.IsNewRow {
			res = append(res, tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(buttonData.Text, strconv.Itoa(buttonData.FuncId))))
		} else {

			if len(res) != 0 {
				lastRowId := len(res) - 1
				res[lastRowId] = append(res[lastRowId], tgbotapi.NewInlineKeyboardButtonData(buttonData.Text, strconv.Itoa(buttonData.FuncId)))
			} else {
				res = append(res, make([]tgbotapi.InlineKeyboardButton, 0))
				res[0] = append(res[0], tgbotapi.NewInlineKeyboardButtonData(buttonData.Text, strconv.Itoa(buttonData.FuncId)))
			}
		}

	}

	return tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: res,
	}
}
