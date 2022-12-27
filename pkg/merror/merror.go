package merror

import "fmt"

type Error interface {
	Error() string
}

type ErrorType struct {
	Status int
	Code   string
	Msg    string
	DevMsg string
}

func (e ErrorType) New() Error {
	return &MError{
		DevMsg: e.DevMsg,
		Msg:    e.Msg,
		Code:   e.Code,
		Status: e.Status,
	}
}

func InternalError() *ErrorType {
	return &ErrorType{
		Status: 500,
		Code:   "E000001",
		Msg:    "Internal Server Error",
		DevMsg: "Internal Server Error",
	}
}

func InvalidParameter() *ErrorType {
	return &ErrorType{
		Status: 400,
		Code:   "E000002",
		Msg:    "Invalid Parameter",
		DevMsg: "Invalid Parameter",
	}
}

type MError struct {
	Status int
	Code   string
	Msg    string
	DevMsg string
}

func (e MError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("%s: %s", e.Code, e.Msg)
	}

	return e.Msg
}

func New(msg string) Error {
	return &MError{
		Status: 500,
		Code:   "",
		Msg:    msg,
		DevMsg: msg,
	}
}
