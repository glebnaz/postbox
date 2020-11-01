package errors

import "fmt"

//Exception custom error type
type Exception struct {
	Message    string
	DevMessage string
	Code       int
}

//New return new Error with err as DevMessage
func (e Exception) New(err error) Exception {
	e.DevMessage = err.Error()
	return e
}

//Error is implement error interface
func (e Exception) Error() string {
	return fmt.Sprintf(`{\n	"code":%v\n		"message":"%v"\n}`, e.Code, e.Message)
}
