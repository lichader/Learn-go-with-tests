package maps

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func Search(dic Dictionary, key string) string {
	return dic[key]
}

func (d Dictionary) Search(key string) (string, error) {
	result, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}

func (d Dictionary) Update(key, value string) error {
	_, ok := d[key]
	if !ok {
		return ErrNotFound
	}

	d[key] = value
	return nil
}
