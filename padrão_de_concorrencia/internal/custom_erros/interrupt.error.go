package customerros

type InterruptError struct{}

func (e *InterruptError) Error() string {
	return "InterruptError"
}
