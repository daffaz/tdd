package main

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type Dictionary map[string]string
type DictionaryErr string

func (d Dictionary) Search(target string) (string, error) {
	val, ok := d[target]
	if ok {
		return val, nil
	}
	return "", ErrorNotFound
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrorNotFound:
		d[key] = value
	case nil:
		return ErrorWordExist
	default:
		return err
	}

	return nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}
