package data

import (
	"time"

	up "github.com/upper/db/v4"
)

type SessionStore struct {
	Session Session
}

func (s *SessionStore) Delete(token string) error {
	return s.Session.DeleteByToken(token)
}

func (s *SessionStore) Find(token string) ([]byte, bool, error) {
	tok, err := s.Session.GetByToken(token)
	if err != nil {
		if err == up.ErrNoMoreRows {
			return nil, false, nil
		}

		return nil, false, err
	}

	return tok.Data, true, nil
}

func (s *SessionStore) Commit(token string, b []byte, expiry time.Time) error {
	s.Session.Token = token
	s.Session.Data = b
	s.Session.Expiry = expiry

	err := s.Session.Upsert()
	if err != nil {
		return err
	}

	return nil
}