package db

import (
	"encoding/json"
	"github.com/NekitMalyarenko/VocabularyBot/types"
)

const (
	nnDataTable         = "nn_data"
	nnDataId            = "id"
	nnDataWord          = "word"
	nnDataWordRating    = "word_rating"
	nnDataDefinitions   = "definitions"
	nnDataUsageExamples = "usage_examples"
	nnDataRatedBy       = "rated_by"
)



func (manager *dbManager) AddNNData(word *types.RowWordData, wordRating float64, userId int64) error {
	nnData, err := parseWordData(word)
	nnData.RatedUserId = userId
	nnData.WordRating = wordRating
	if err != nil {
		return err
	}

	_, err = manager.Session.InsertInto(nnDataTable).
		Values(nnData).
		Exec()

	return err
}


func (manager *dbManager) HasNNWord(word string) bool {
	query := manager.Session.SelectFrom(nnDataTable)
	query = query.Where(nnDataWord + "='" + word + "'")

	return query.Iterator().Next()
}


func parseWordData(word *types.RowWordData) (*types.NNData, error) {
	data := new(types.NNData)
	data.Word = word.Word

	temp, err := json.Marshal(word.Definitions)
	if err != nil {
		return new(types.NNData), err
	}
	data.Definitions = string(temp)

	temp, err = json.Marshal(word.UsageExamples)
	if err != nil {
		return new(types.NNData), err
	}
	data.UsageExamples = string(temp)

	data.WordRank = word.WordRank

	return data, nil
}