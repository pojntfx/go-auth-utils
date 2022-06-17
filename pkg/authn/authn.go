package authn

import (
	"context"
	"errors"
)

var (
	ErrWrongUsername = errors.New("wrong username")
	ErrWrongPassword = errors.New("wrong password")
)

type Authn interface {
	Open(context.Context) error
	Validate(username, token string) error
}
