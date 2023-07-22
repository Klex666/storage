package storage

import "errors"

var store = make(map[string]string)
var ErrorNoSuchKey = errors.New("no such key")

func Put(key string, value string) error {
	store[key] = value

	return nil
}

func Get(key string) (string, error) {
	value, err := store[key]
	if !err {
		return "", ErrorNoSuchKey
	}
	return value, nil
}

func Delete(key string) error {
	delete(store, key)

	return nil
}
