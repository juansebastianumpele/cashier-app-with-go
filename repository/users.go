package repository

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) ReadUser() ([]model.Credentials, error) {
	records, err := u.db.Load("users")
	if err != nil {
		return nil, err
	}

	var listUser []model.Credentials
	err = json.Unmarshal([]byte(records), &listUser)
	if err != nil {
		return nil, err
	}

	return listUser, nil
}

func (u *UserRepository) AddUser(creds model.Credentials) error {
	model, err := u.ReadUser()
	if err != nil {
		return err
	}
	for _, v := range model {
		creds.Username = v.Username
		creds.Password = v.Password

		// if v.Username == creds.Username {

		// }

	}
	model = append(model, creds)
	data, err := json.Marshal(model)

	if err != nil {
		return err
	}

	err = u.db.Save("users", data)
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (u *UserRepository) ResetUser() error {
	err := u.db.Reset("users", []byte("[]"))
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) LoginValid(req model.Credentials) error {
	listUser, err := u.ReadUser()
	if err != nil {
		return err
	}

	for _, element := range listUser {
		if element.Username == req.Username && element.Password == req.Password {
			return nil
		}
	}

	return fmt.Errorf("Wrong User or Password!")
}
