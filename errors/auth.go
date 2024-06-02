package errors


import (
	"fmt"
)

var (
	ErrInvalidToken        = fmt.Errorf("invalid token")                  //AuthError
	ErrInactiveUser        = fmt.Errorf("inactive user")                  //AuthError
	ErrInvalidUser         = fmt.Errorf("invalid user")                   //AuthError        //AuthError
)