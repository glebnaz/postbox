package errors

import (
	"fmt"
	"log"
)

//Exception custom error type
type Exception struct {
	Message    string
	DevMessage string
	Code       int
}

//New return new Error with err as DevMessage
//and log dev error by log
func (e Exception) New(err error) Exception {
	log.Printf("Code: %v;\n Message: %v\n", e.Code, err)
	e.DevMessage = err.Error()
	return e
}

//Error is implement error interface
func (e Exception) Error() string {
	return fmt.Sprintf(`{\n	"code":%v\n		"message":"%v"\n}`, e.Code, e.Message)
}
