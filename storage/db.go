package storage

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type DB interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte) error
}