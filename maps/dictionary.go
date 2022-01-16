package main

type Dictionary map[string]string

func (d Dictionary) Search(target string) string {
	return d[target]
}
