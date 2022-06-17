package basic

import (
	"context"

	"github.com/pojntfx/go-auth-utils/pkg/authn"
)

type Authn struct {
	username string
	password string
}

func NewAuthn(username, password string) *Authn {
	return &Authn{
		username: username,
		password: password,
	}
}

func (a *Authn) Open(context.Context) error {
	return nil
}

func (a *Authn) Validate(username, token string) error {
	if username != a.username {
		return authn.ErrWrongUsername
	}

	if token != a.password {
		return authn.ErrWrongPassword
	}

	return nil
}
