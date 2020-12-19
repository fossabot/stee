package stee

import (
	"errors"
)

type redirections map[string]string

var (
	// ErrRedirectionNotfound is used when the redirection was not found
	ErrRedirectionNotfound = errors.New("No redirection found for this key")
	// ErrRedirectionAlreadyExists is used when the redirection already exists
	ErrRedirectionAlreadyExists = errors.New("A redirection is already associated with this key")
	// ErrTargetIsNotAValidURL is used when the provided target is not a valid URL
	ErrTargetIsNotAValidURL = errors.New("The redirection target is not a valid URL")
)

// GetRedirection gets a redirection based on its key
func (c *Core) GetRedirection(key string) (target string, err error) {
	target, err = c.store.ReadRedirection(key)
	return
}

// AddRedirection adds a redirection. It takes both the key and the target of the redirection.
func (c *Core) AddRedirection(key string, target string) (err error) {
	err = c.store.WriteRedirection(key, target)
	return
}

// DeleteRedirection deletes a redirection based on its key.
func (c *Core) DeleteRedirection(key string) (err error) {
	err = c.store.DeleteRedirection(key)
	return err
}
