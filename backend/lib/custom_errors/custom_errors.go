package custom_errors

import (
	"fmt"
	"runtime/debug"
)

const (
	DEFAULT_ERROR_CODE        = 0
	GITHUB_API_40X_ERROR_CODE = 100
	GITHUB_API_50X_ERROR_CODE = 200
)

type CustomError struct {
	Inner      error
	Message    string
	Code       int
	StackTrace string
	Data       map[string]any
}

func (err CustomError) Error() string {
	return err.Message
}

func Wrap(err error, format string, args ...any) CustomError {
	return CustomError{
		Inner:      err,
		Message:    fmt.Sprintf(format, args...),
		Code:       DEFAULT_ERROR_CODE,
		StackTrace: string(debug.Stack()),
		Data:       make(map[string]any),
	}
}
