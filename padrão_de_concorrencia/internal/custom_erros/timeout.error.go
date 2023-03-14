package customerros

type TimeoutError struct {
	message string
}

func (e *TimeoutError) Error() string {
	return e.message
}
