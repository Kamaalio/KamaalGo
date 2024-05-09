package strings

import (
	"errors"
)

var ErrStringUnwrap = errors.New("value not defined")

func Unwrap(value string) (string, error) {
	if value == "" {
		return "", ErrStringUnwrap
	}

	return value, nil
}
