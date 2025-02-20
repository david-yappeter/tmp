package util

type AppError struct {
	Message string
	Code    int
}

func (e *AppError) Error() string {
	return e.Message
}

func NewBadRequestError(msg string) error {
	return &AppError{Message: msg, Code: 400}
}

func NewInternalServerError(msg string) error {
	return &AppError{Message: msg, Code: 500}
}
