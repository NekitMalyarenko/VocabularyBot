package web

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"vocabulary-tg-bot/db"
	"vocabulary-tg-bot/web/types"
)

const (
	WORD_COUNT_URL        = "http://www.wordcount.org/dbquery.php?"
	OXFORD_DICTIONARY_URL = "https://en.oxforddictionaries.com/definition/"
	NUMBER_OF_EXAMPLES    = 5
)

func GetNNTrainingWord() *webTypes.RowWordData {

	for {
		word, err := GetRandomWord()
		if err != nil {
			log.Println(err)
			continue
		}

		if !db.GetDBManager().HasNNWord(word.Word) {
			return word
		}
	}

}

func GetRandomWord() (*webTypes.RowWordData, error) {
	var (
		err      error
		wordData = new(webTypes.RowWordData)
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

		if checkData(wordData) {
			return wordData, nil
		} else {
			log.Println("failed =6", wordData.WordRank)
		}
	}
}

func getWordData(index int) (word string, err error) {
	url := WORD_COUNT_URL + "toFind=" + strconv.Itoa(index) + "&method=SEARCH_BY_INDEX"

	resp, err := new(http.Client).Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	parsedResponse := buffer.String()
	word = parsedResponse[strings.Index(parsedResponse, "word0")+6 : strings.Index(parsedResponse, "&freq0")]

	return word, nil
}

func getExtendedWordData(word string) (Definitions map[string][]string, UsageExamples map[string][]string, err error) {
	urlToParse := OXFORD_DICTIONARY_URL + word
	wordPage, err := goquery.NewDocument(urlToParse)
	if err != nil {
		return nil, nil, err
	}

	definitions := make(map[string][]string)
	usageExamples := make(map[string][]string)

	wordPage.Find(".gramb").Each(func(i int, s *goquery.Selection) {
		languagePart := s.Find("h3 .pos").Text()

		if languagePart != "" {
			usageExample := s.Find(".semb li .trg .examples .exg .ex em").Text()
			usageExamples[languagePart] = parseUsageExamples(usageExample)

			defNodes := s.Find(".semb .trg .ind").Nodes
			//log.Println(defNodes)
			for _, val := range defNodes {
				//log.Println(val.Attr)
				definitions[languagePart] = append(definitions[languagePart], s.FindNodes(val).Text())
			}
		}
	})
	return definitions, usageExamples, nil
}

func parseUsageExamples(input string) []string {
	var (
		startIndex int
		endIndex   int
		result     []string
	)
	result = make([]string, 0)

	input = strings.Replace(input, "\n", "", -1)
	input = strings.Replace(input, "‘‘", "‘", -1)
	input = strings.Replace(input, "’’", "’", -1)
	input = strings.Replace(input, "’‘", "^_", -1)
	input = strings.Replace(input, "‘", "_", 1)

	if len(input) == 0 {
		return result
	}

	temp := []byte(input)
	temp[len(temp)-1] = '^'
	input = string(temp)

	//log.Println(input)

	for len(input) > 0 && len(result) < NUMBER_OF_EXAMPLES {
		startIndex = strings.Index(input, "_") + 1
		endIndex = strings.Index(input, "^")

		//log.Println(startIndex, ":", endIndex)
		//log.Println("cut part:", input[startIndex:endIndex])

		result = append(result, input[startIndex:endIndex])
		input = input[endIndex+1:]
		//log.Println("after cut:", input)
	}

	return result
}

func checkData(data *webTypes.RowWordData) bool {

	if len(data.Definitions) == 0 || len(data.UsageExamples) == 0 {
		return false
	}

	for _, val := range data.Definitions {
		//log.Println("definitoons of", key, "len:", len(val))
		if len(val) == 0 {
			return false
		}

		if len(val) == 1 && (val[0] == "" || val[0] == " ") {
			return false
		}
	}

	for _, val := range data.UsageExamples {
		//log.Println("usage examples of", key, "len:", len(val))
		if len(val) == 0 {
			return false
		}

		if len(val) == 1 && (val[0] == "" || val[0] == " ") {
			return false
		}
	}

	return true
}
