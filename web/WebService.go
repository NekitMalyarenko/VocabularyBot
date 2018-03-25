package web

import (
	"strings"
	"github.com/NekitMalyarenko/VocabularyBot/types"
	"github.com/NekitMalyarenko/VocabularyBot/bot/services"
)


func ToString(word *types.RowWordData) string {
	text := telegramServices.MessageBuilderInit().BoldText(strings.ToUpper(word.Word)).Text

	for key, value := range word.Definitions {
		text += telegramServices.MessageBuilderInit().NewRow().ItalicText(key).NormalText("\n" + value[0]).Text
		text += telegramServices.MessageBuilderInit().NormalText("\n\n").CodeText(word.UsageExamples[key][0]).Text
	}

	return text
}
