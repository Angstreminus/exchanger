package apperrors

type Apperror interface {
	Error() string
}

// Custom error type for invalid data fomat (negative/non-int values etc)
type InvalidData struct {
	Message string
}

func (ae InvalidData) Error() string {
	return ae.Message
}
