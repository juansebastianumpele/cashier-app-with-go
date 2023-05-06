package repository

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SessionsRepository struct {
	db db.DB
}

func NewSessionsRepository(db db.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) ReadSessions() ([]model.Session, error) {
	records, err := u.db.Load("sessions")
	if err != nil {
		return nil, err
	}

	var listSessions []model.Session
	err = json.Unmarshal([]byte(records), &listSessions)
	if err != nil {
		return nil, err
	}

	return listSessions, nil
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return err
	}

	for i, v := range listSessions {
		if tokenTarget == v.Token {
			listSessions = append(listSessions[:i], listSessions[:i]...)
		}
	}

	// Select target token and delete from listSessions
	// TODO: answer here

	jsonData, err := json.Marshal(listSessions)
	if err != nil {
		return err
	}

	err = u.db.Save("sessions", jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	model, err := u.ReadSessions()
	if err != nil {
		return err
	}
	// for _, v := range model {
	// 	// v.Token = uuid.New().String()
	// 	session.Token = v.Token
	// 	session.Username = v.Username
	// 	session.Expiry = v.Expiry
	// }
	model = append(model, session)

	data, _ := json.Marshal(model)
	err = u.db.Save("sessions", data)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) CheckExpireToken(token string) (model.Session, error) {
	isExist, err := u.TokenExist(token)
	if err != nil {
		return model.Session{}, err
	}
	Expired := u.TokenExpired(isExist)

	if Expired {
		u.DeleteSessions(isExist.Token)
		return model.Session{}, errors.New("Token is Expired!")
	}

	return isExist, nil // TODO: replace this
}

func (u *SessionsRepository) ResetSessions() error {
	err := u.db.Reset("sessions", []byte("[]"))
	if err != nil {
		return err
	}
	return nil
}

func (u *SessionsRepository) TokenExist(req string) (model.Session, error) {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return model.Session{}, err
	}
	for _, element := range listSessions {
		if element.Token == req {
			return element, nil
		}
	}
	return model.Session{}, fmt.Errorf("Token Not Found!")
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
