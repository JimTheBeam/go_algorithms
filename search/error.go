package search

import "errors"

var ErrNotFound error = errors.New("Not found")

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string {
	return e.Name + "not found"
}
