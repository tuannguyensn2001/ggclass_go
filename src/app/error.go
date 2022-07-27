package app

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpError struct {
	StatusCode int         `json:"statusCode" yaml:"statusCode"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message" yaml:"message"`
	RootErr    error       `json:"-"`
	Code       string      `json:"code" yaml:"code"`
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

func LoadErr(url string) (map[string]map[string]HttpError, error) {
	yamlFile, err := ioutil.ReadFile(url)
	if err != nil {
		return nil, err
	}

	m := make(map[string]map[string]HttpError)

	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func ParseError(key string) *HttpError {
	m, err := LoadErr("../../error.yml")

	if err != nil {
		return InternalHttpError("server has error", errors.New("error"))
	}

	split := strings.Split(key, ".")
	if len(split) != 2 {
		return InternalHttpError("server has error", errors.New("error"))
	}
	first := split[0]
	second := split[1]

	val, ok := m[first]
	if !ok {
		return InternalHttpError("server has error", errors.New("error"))
	}
	check, ok := val[second]
	if !ok {
		return InternalHttpError("server has error", errors.New("error"))
	}

	return &check

}
