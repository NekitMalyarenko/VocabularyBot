package db

import (
	"log"
	"upper.io/db.v3"
	"github.com/NekitMalyarenko/VocabularyBot/types"
)

const (
	usersTable     = "users"
	usersId        = "id"
	usersFirstName = "first_name"
	usersLastName  = "last_name"
	usersUserName  = "user_name"
	usersIsTester  = "is_tester"
)


func (manager *dbManager) GetAllUsers() ([]*types.User, error) {
	var users []*types.User

	res := manager.Session.Collection(usersTable).Find()
	err := res.All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}


func (manager *dbManager) GetUser(id int64) (*types.User, error) {
	var user *types.User

	res := manager.Session.Collection(usersTable).Find(db.Cond{usersId: id})
	err := res.All(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}


func (manager *dbManager) GetAllTesters() ([]*types.User, error) {
	var users []*types.User

	res := manager.Session.Collection(usersTable).Find(db.Cond{usersIsTester: true})
	err := res.All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}


func (manager *dbManager) HasUser(id int64) (bool, error) {
	var users []*types.User

	res := manager.Session.Collection(usersTable).Find(db.Cond{usersId: id})
	err := res.All(&users)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return len(users) != 0, nil
}


func (manager *dbManager) IsUserTester(id int64) (bool, error) {
	var users []*types.User

	res := manager.Session.Collection(usersTable).Find(db.Cond{usersId: id, usersIsTester: true})
	err := res.All(&users)
	if err != nil {
		return false, err
	}

	return len(users) != 0, nil
}


func (manager *dbManager) AddUser(user types.User) error {
	_, err := manager.Session.InsertInto(usersTable).
		Values(user).
		Exec()

	return err
}