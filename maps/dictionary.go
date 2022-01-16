package main

import (
	"errors"
)

var ErrorNotFound = errors.New("not found")

type Dictionary map[string]string

func (d Dictionary) Search(target string) (string, error) {
	val, ok := d[target]
	if ok {
		return val, nil
	}
	return "", ErrorNotFound
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
