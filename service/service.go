package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"

	"zt_test/entity"
)

type Storage interface {
	AddUser(user entity.User) error
	GetLastUser() (entity.User, error)
	IncreaseCache(ctx context.Context, key string, val int64) (int64, error)
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{storage: s}
}

func (s *Service) AddNewUser(user entity.User) error {
	err := s.storage.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetUser() (entity.User, error) {
	user, err := s.storage.GetLastUser()
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (s *Service) ComputeHmac(text, key string) (string, error) {
	keydump := hex.Dump([]byte(key))
	h := hmac.New(sha512.New, []byte(keydump))
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil)), nil
}

func (s *Service) Increase(ctx context.Context, key string, value int64) (int64, error) {
	res, err := s.storage.IncreaseCache(ctx, key, value)
	if err != nil {
		return 0, err
	}
	return res, nil
}
