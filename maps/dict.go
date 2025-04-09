

package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound = errors.New("could not find the word that you look for")
	ErrWordExists = errors.New("cannot add word bcz it already exists")
	ErrWordDoesNotExist = errors.New("cannot update/delete that word does not exists")
)



func (d Dictionary) Search(word string) (string, error) {

	defination, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return defination, nil
}

func (d Dictionary) Add(word, defination string) error {

	_, err := d.Search(word)

	switch err {

	case ErrNotFound:
		d[word] = defination

	case nil:
		return ErrWordExists

	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefination string) error {

	_, err := d.Search(word)

	switch err {

	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefination
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	if err != nil {
		return ErrWordDoesNotExist
	}

	delete(d, word)

	return nil
}

