package kv

import (
	"context"
	"errors"
	"sync"
)

type Storage struct {
	m sync.Map
}

func NewLinkKVStorage() *Storage {
	return &Storage{m: sync.Map{}}
}

func (s *Storage) Insert(_ context.Context, key string, value []byte) error {
	s.m.Store(key, value)
	return nil
}

func (s *Storage) GetByID(_ context.Context, key string) ([]byte, error) {
	v, in := s.m.Load(key)
	if !in {
		return nil, errors.New("not found")
	}

	return v.([]byte), nil
}
