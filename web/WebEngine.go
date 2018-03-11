package web

import (
	"math/rand"
	"strconv"
	"strings"
	"net/http"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"time"
	"telegram/helpers"
	"log"
)


type RowWordData struct {
	Word           string
	Definitions    map[string][]string
	UsageExamples  map[string][]string
	WordRank       int

}


const (
	WORD_COUNT_URL        = "http://www.wordcount.org/dbquery.php?"
	OXFORD_DICTIONARY_URL = "https://en.oxforddictionaries.com/definition/"
	NUMBER_OF_EXAMPLES    = 1
)



func GetRandomWord() (*RowWordData, error) {
	var (
		err error
		wordData = new(RowWordData)
	)
	rand.Seed(time.Now().UnixNano())

	for {
		wordData.WordRank = rand.Intn(86799)

		wordData.Word, err = getWordData(wordData.WordRank)
		if err != nil {
			return nil, err
		}

		wordData.Definitions, wordData.UsageExamples, err = getExtendedWordData(wordData.Word)
		if err != nil {
			return nil, err
		}

		if len(wordData.Definitions) != 0 && len(wordData.UsageExamples) != 0 {
			return wordData, nil
		}
	}
}


func getWordData(index int) (word string, err error)  {
	url := WORD_COUNT_URL + "toFind=" + strconv.Itoa(index) + "&method=SEARCH_BY_INDEX"

	resp, err := new(http.Client).Get(url)
	if err != nil {
		return "",  err
	}
	defer resp.Body.Close()

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	parsedResponse := buffer.String()
	word = parsedResponse[strings.Index(parsedResponse, "word0") + 6 : strings.Index(parsedResponse, "&freq0")]

	return word,nil
}


func getExtendedWordData(word string) (Definitions map[string][]string, UsageExamples map[string][]string, err error) {
	urlToParse := OXFORD_DICTIONARY_URL + word
	wordPage,err := goquery.NewDocument(urlToParse)
	if err!=nil {
		return nil,nil,err
	}

	definitions := make(map[string][]string)
	usageExamples := make( map[string][]string )

	wordPage.Find(".gramb ").Each(func(i int, s *goquery.Selection) {
		languagePart := s.Find("h3 .pos").Text()

		if len(usageExamples[languagePart]) < NUMBER_OF_EXAMPLES {
			usageExample := s.Find(".semb li .trg .examples .exg .ex em").Text()
			usageExamples[languagePart] = append(usageExamples[languagePart], cutUsageExamples(usageExample))
		}

		definition := s.Find(".semb .trg .ind").Text()
		definitions[languagePart] = append(definitions[languagePart] , definition)
	})
	return definitions, usageExamples,nil
}


func cutUsageExamples(input string) string {

	if strings.Count(input , "‘")	> NUMBER_OF_EXAMPLES {
		input = strings.Replace(input, "‘", "{", 1)
		input = strings.Replace(input, "’", "}", strings.Count(input , "‘") - NUMBER_OF_EXAMPLES)

		log.Println(input)

		return input[strings.LastIndex(input, "}") + 1:]
	} else {
		return input
	}
}


func (word *RowWordData) ToString() string {
	text := telegramHelpers.MessageBuilderInit().BoldText(word.Word).Text

	for key, value := range word.Definitions {
		text += telegramHelpers.MessageBuilderInit().NewRow().ItalicText(key).Text
		for _, definition := range value {
			text += telegramHelpers.MessageBuilderInit().NormalText("\n" + definition).Text
		}
	}

	return text
}