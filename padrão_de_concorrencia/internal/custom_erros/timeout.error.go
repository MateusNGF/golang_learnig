package customerros

import "errors"

type TimeoutError interface {
	error
}

func NewTimeoutError() error {
	return errors.New("Timeout")
}
