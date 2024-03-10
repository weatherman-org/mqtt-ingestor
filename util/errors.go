package util

import (
	"errors"
	"net/http"
)

var (
	ErrDatabase                      = errors.New("database error")
	ErrUnableToDecodeJSON            = errors.New("unable to decode json, input format is wrong")
	ErrInputJSONMustOnlyHaveOneValue = errors.New("input json must only have one value")
)

var CustomErrorType = map[error]int{
	ErrDatabase: http.StatusInternalServerError,
}
