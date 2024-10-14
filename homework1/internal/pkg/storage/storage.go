package storage

import (
	"go.uber.org/zap"
	"strconv"
)

type Storage struct {
	storage map[string]Value
	logger  *zap.Logger
}

type Value struct {
	v    string
	kind string
}

func InitStorage() (Storage, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return Storage{}, err
	}
	defer logger.Sync()
	logger.Info("storage initialized")
	return Storage{
		storage: make(map[string]Value),
		logger:  logger,
	}, nil
}

func (s *Storage) Set(key, value string) {
	var val Value
	val.v = value
	if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		val.kind = "D"
	} else {
		val.kind = "S"
	}
	s.storage[key] = val
	s.logger.Info("storage set key-value")
}

func (s *Storage) Get(key string) *string {
	out, ok := s.storage[key]
	if !ok {
		return nil
	}
	s.logger.Info("storage get key")
	return &out.v
}

func (s *Storage) GetKind(key string) string {
	out, ok := s.storage[key]
	s.logger.Info("Get kind key value")
	if !ok {
		return "N"
	}
	return out.kind
}
