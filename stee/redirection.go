package stee

import (
	"errors"
)

type redirections map[string]string

// Errors about redirections
var (
	ErrRedirectionNotfound      = errors.New("No redirection found for this key")
	ErrRedirectionAlreadyExists = errors.New("A redirection is already associated with this key")
	ErrTargetIsNotAValidURL     = errors.New("The redirection target is not a valid URL")
)

func (c *Core) GetRedirection(key string) (target string, ok bool) {
	target, ok = (*c.store).ReadRedirection(key)
	return
}

func (c *Core) AddRedirection(key string, target string) (err error) {
	err = (*c.store).WriteRedirection(key, target)
	return
}

func (c *Core) DeleteRedirection(key string) (err error) {
	err = (*c.store).DeleteRedirection(key)
	return err
}
