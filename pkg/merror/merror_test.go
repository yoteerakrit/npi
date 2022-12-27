package merror

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	err := InternalError().New()
	assert.Equal(t, &MError{
		Status: 500,
		Code:   "E000001",
		Msg:    "Internal Server Error",
		DevMsg: "Internal Server Error",
	}, err)

	err1 := &ErrorType{
		Status: 400,
		Code:   "Test",
		Msg:    "Test",
		DevMsg: "Test",
	}

	assert.Equal(t, &MError{
		Status: 400,
		Code:   "Test",
		Msg:    "Test",
		DevMsg: "Test",
	}, err1.New())

}
