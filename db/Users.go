package db

import (
	"log"
	"upper.io/db.v3"
	"github.com/NekitMalyarenko/VocabularyBot/types"
)

const (
	USERS_TABLE      = "users"
	USERS_ID         = "id"
	USERS_FIRST_NAME = "first_name"
	USERS_LAST_NAME  = "last_name"
	USERS_USER_NAME  = "user_name"
	USERS_IS_TESTER  = "is_tester"
)


func (manager *dbManager) GetAllUsers() ([]*types.User, error) {
	var users []*types.User

	res := manager.Session.Collection(USERS_TABLE).Find()
	err := res.All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}


func (manager *dbManager) GetUser(id int64) (*types.User, error) {
	var user *types.User

	res := manager.Session.Collection(USERS_TABLE).Find(db.Cond{USERS_ID: id})
	err := res.All(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}


func (manager *dbManager) GetAllTesters() ([]*types.User, error) {
	var users []*types.User

	res := manager.Session.Collection(USERS_TABLE).Find(db.Cond{USERS_IS_TESTER: true})
	err := res.All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}


func (manager *dbManager) HasUser(id int64) (bool, error) {
	var users []*types.User

	res := manager.Session.Collection(USERS_TABLE).Find(db.Cond{USERS_ID: id})
	err := res.All(&users)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return len(users) != 0, nil
}


func (manager *dbManager) IsUserTester(id int64) (bool, error) {
	var users []*types.User

	res := manager.Session.Collection(USERS_TABLE).Find(db.Cond{USERS_ID: id, USERS_IS_TESTER: true})
	err := res.All(&users)
	if err != nil {
		return false, err
	}

	return len(users) != 0, nil
}


func (manager *dbManager) AddUser(user types.User) error {
	_, err := manager.Session.InsertInto(USERS_TABLE).
		Values(user).
		Exec()

	return err
}