package customerros

import "errors"

type InterruptError interface {
	error
}

func NewInterruptError() error {
	return errors.New("Interrupteded.")
}
