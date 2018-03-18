package db

import (
	"encoding/json"
	"vocabulary-tg-bot/web/types"
)

const (
	NN_DATA_TABLE          = "nn_data"
	NN_DATA_ID             = "id"
	NN_DATA_WORD           = "word"
	NN_DATA_WORD_RATING    = "word_rating"
	NN_DATA_DEFINITIONS    = "definitions"
	NN_DATA_USAGE_EXAMPLES = "usage_examples"
	NN_DATA_RATED_BY       = "rated_by"
)

type NNData struct {
	Id            int64   `db:"id,omitempty"`
	Word          string  `db:"word"`
	WordRank      int     `db:"word_rank"`
	Definitions   string  `db:"definitions"`
	UsageExamples string  `db:"usage_examples"`
	WordRating    float64 `db:"word_rating"`
	RatedUserId   int64   `db:"rated_by"`
}

func (manager *dbManager) AddNNData(word *webTypes.RowWordData, wordRating float64, userId int64) error {
	nnData, err := new(NNData).parseWordData(word)
	nnData.RatedUserId = userId
	nnData.WordRating = wordRating
	if err != nil {
		return err
	}

	_, err = manager.Session.InsertInto(NN_DATA_TABLE).
		Values(nnData).
		Exec()

	return err
}

func (manager *dbManager) HasNNWord(word string) bool {
	query := manager.Session.SelectFrom(NN_DATA_TABLE)
	query = query.Where(NN_DATA_WORD + "='" + word + "'")

	return query.Iterator().Next()
}

func (data *NNData) parseWordData(word *webTypes.RowWordData) (*NNData, error) {
	data.Word = word.Word

	temp, err := json.Marshal(word.Definitions)
	if err != nil {
		return new(NNData), err
	}
	data.Definitions = string(temp)

	temp, err = json.Marshal(word.UsageExamples)
	if err != nil {
		return new(NNData), err
	}
	data.UsageExamples = string(temp)

	data.WordRank = word.WordRank

	return data, nil
}
