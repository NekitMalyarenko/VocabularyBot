package db

import (
	"github.com/NekitMalyarenko/VocabularyBot/types"
	"upper.io/db.v3"
	"log"
)

const (
	wordsTable         = "words"
	wordsId            = "id"
	wordsDate          = "date"
	wordsWord          = "word"
	wordsDefinitions   = "definitions"
	wordsUsageExamples = "usage_examples"
	wordsLikes         = "likes"
	wordsDislikes      = "dislikes"
)


func (manager *dbManager) AddWord(word *types.Word) error {
	_, err := manager.Session.InsertInto(wordsTable).Values(word).Exec()
	return err
}


func (manager *dbManager) GetWord(date string) (word *types.Word, err error) {
	res := manager.Session.Collection(wordsTable).Find(db.Cond{wordsDate: date})
	err = res.One(&word)
	if err != nil {
		return nil, err
	}

	return word, err
}


func (manager *dbManager) HasWord(word string) bool {
	var words []*types.Word
	res := manager.Session.Collection(wordsTable).Find(db.Cond{wordsWord: word})
	err := res.One(&words)
	if err != nil {
		log.Println(err)
		return false
	}

	return len(words) != 0
}