package app

import "net/http"

type HttpError struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	RootErr    error       `json:"-"`
}

func BadRequestHttpError(message string, err error) *HttpError {
	return &HttpError{
		StatusCode: http.StatusBadRequest,
		Data:       nil,
		RootErr:    err,
		Message:    message,
	}
}

func (err *HttpError) Error() string {
	return err.Message
}

func InternalHttpError(message string, err error) *HttpError {
	return &HttpError{
		StatusCode: http.StatusInternalServerError,
		RootErr:    err,
		Message:    message,
	}
}

func ConflictHttpError(message string, err error) *HttpError {
	return &HttpError{
		StatusCode: http.StatusConflict,
		RootErr:    err,
		Message:    message,
	}
}

func ForbiddenHttpError(message string, err error) *HttpError {
	return &HttpError{
		StatusCode: http.StatusForbidden,
		RootErr:    err,
		Message:    message,
	}
}

func NotFoundHttpError(message string, err error) *HttpError {
	return &HttpError{
		StatusCode: http.StatusNotFound,
		RootErr:    err,
		Message:    message,
	}
}
