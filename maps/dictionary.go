package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordAlreadyExists = DictionaryErr("key already exists")
	ErrWordDoesNotExist  = DictionaryErr("Word does not exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordAlreadyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
