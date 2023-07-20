package errors

import (
	"fmt"
	"runtime/debug"
)

const (
	DEFAULT_ERROR_CODE        = 0
	GITHUB_API_40X_ERROR_CODE = 100
	GITHUB_API_50X_ERROR_CODE = 200
)

func Test() {
	err := Test1()
	customErr, _ := err.(CustomError)
	fmt.Println(customErr.HasCode(GITHUB_API_50X_ERROR_CODE))
}

func Test1() error {
	err := Test2()
	return Wrap(err, "Failed to call Test2")
}

func Test2() error {
	err := Lib()
	customErr := Wrap(err, "Failed to call Lib")
	customErr.Code = GITHUB_API_40X_ERROR_CODE
	return customErr
}

func Lib() error {
	return fmt.Errorf("Error from Lib")
}

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

func (err CustomError) HasCode(code int) bool {
	if err.Code == code {
		return true
	}
	if err.Inner != nil {
		inner, ok := err.Inner.(CustomError)
		if !ok {
			return false
		}
		return inner.HasCode(code)
	}
	return true
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
