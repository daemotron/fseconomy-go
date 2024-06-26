package security

import (
	"github.com/fseconomy/fseconomy-go/internal/security"
)

type Security struct {
	security.Security
}

// WithSecretKey sets the secret key
func WithSecretKey(secretKey string) func(*Security) error {
	return func(a *Security) error {
		return a.SetSecretKeyFromString(secretKey)
	}
}

// WithSecretIv sets the initialization vector
func WithSecretIv(secretIv string) func(*Security) error {
	return func(a *Security) error {
		return a.SetSecretIvFromString(secretIv)
	}
}

// WithUser sets the user name
func WithUser(username string) func(*Security) error {
	return func(a *Security) error {
		a.SetUserName(username)
		return nil
	}
}

// WithToken sets the auth token
func WithToken(token string) func(*Security) error {
	return func(a *Security) error {
		a.SetAuthToken(token)
		return nil
	}
}

// New returns a new Fseconomy security object
func New(options ...func(*Security) error) (*Security, error) {
	sec := &Security{}
	err := sec.SetSecretKeyFromEnv()
	if err != nil {
		return nil, err
	}
	err = sec.SetSecretIvFromEnv()
	if err != nil {
		return nil, err
	}
	for _, option := range options {
		if err := option(sec); err != nil {
			return nil, err
		}
	}
	return sec, nil
}
