package maps

import "errors"

var ErrNotFound = errors.New("could not find that word")

type Dict map[string]string

func (d Dict) Definition(lookup string) (string, error) {
	definition, ok := d[lookup]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dict) Add(word, definition string) {
	d[word] = definition
}
