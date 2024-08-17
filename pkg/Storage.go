package goevm

import "github.com/ethereum/go-ethereum/common"

type Storage struct {
	data  map[int]common.Hash //Q
	cache map[int]bool
}

func (s *Storage) Load(key int) (value common.Hash, isWarm bool) {
	isWarm = s.cache[key]

	if !isWarm {
		s.cache[key] = true
	}
	value, ok := s.data[key]
	if !ok {
		return common.Hash{}, false
	}
	return value, ok
}

func (s *Storage) Store(key int, value common.Hash) (isWarm bool) {
	isWarm = s.cache[key]
	if !isWarm {
		s.cache[key] = true
	}
	s.data[key] = value
	return isWarm
}

func NewStorage() *Storage {
	return &Storage{
		data:  make(map[int]common.Hash),
		cache: make(map[int]bool),
	}
}
