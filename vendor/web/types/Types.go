package webTypes

import (
	"strings"
	"telegram/helpers"
)

type RowWordData struct {
	Word          string
	Definitions   map[string][]string
	UsageExamples map[string][]string
	WordRank      int
}

func (word *RowWordData) ToString() string {
	text := telegramHelpers.MessageBuilderInit().BoldText(word.Word).Text

	for key, value := range word.Definitions {
		text += telegramHelpers.MessageBuilderInit().NewRow().BoldText(strings.ToUpper(key)).NormalText("\n" + value[0]).Text
		text += telegramHelpers.MessageBuilderInit().NormalText("\n\n").CodeText(word.UsageExamples[key][0]).Text
	}

	return text
}
