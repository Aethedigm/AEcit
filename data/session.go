package data

import (
	"errors"
	"time"

	up "github.com/upper/db/v4"
)

type Session struct {
	Token  string    `db:"token"`
	Data   []byte    `db:"data"`
	Expiry time.Time `db:"expiry"`
}

func (s *Session) Table() string {
	return "sessions"
}

func (s *Session) GetByToken(token string) (*Session, error) {
	collection := upper.Collection(s.Table())

	var newSession Session
	res := collection.Find(up.Cond{"token": token})

	if err := res.One(&newSession); err != nil {
		return nil, err
	}

	return &newSession, nil
}

func (s *Session) DeleteByToken(token string) error {
	collection := upper.Collection(s.Table())

	res := collection.Find(up.Cond{"token": token})
	err := res.Delete()
	if err != nil {
		return errors.New("error deleting session")
	}

	return nil
}

func (s *Session) Insert() error {
	collection := upper.Collection(s.Table())

	_, err := collection.Insert(s)
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) Update() error {
	collection := upper.Collection(s.Table())

	res := collection.Find(up.Cond{"token": s.Token})
	err := res.Update(s)
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) Upsert() error {
	_, err := s.GetByToken(s.Token)
	if err != nil {
		// Session not found, insert
		err = s.Insert()
		if err != nil {
			return err
		}

		return nil
	}

	// Session found, update
	err = s.Update()
	if err != nil {
		return err
	}

	return nil
}
