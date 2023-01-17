package dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotfound = errors.New("Not found")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotfound

}
