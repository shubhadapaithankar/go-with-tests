package main

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound         = errors.New("could not find the word you were looking for")
	ErrWordExists       = errors.New("Unable to add word because it already exists")
	ErrWordDoesNotExist = errors.New("Unble to update word because it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, exists := d[word]
	if !exists {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	if err == ErrNotFound {
		d[word] = definition
		return nil
	} else if err == nil {
		return ErrWordExists
	}

	return err
}

func (d Dictionary) Update(word, definition string) error {
	_, exists := d[word]
	if exists {
		d[word] = definition
		return nil
	}
	return ErrWordDoesNotExist
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
