package stee

import (
	"fmt"

	"github.com/milanrodriguez/stee/internal/storage"
	"github.com/spf13/viper"
)

// Core is the central manager of Stee. Primary interaction point
type Core struct {
	store storage.Store
}

type coreOption func(*Core) error

// NewCore returns a new core.
func NewCore(options ...coreOption) (*Core, error) {
	var err error
	c := &Core{}

	for _, option := range options {
		err = option(c)
		if err != nil {
			return c, err
		}
	}

	if c.store == nil {
		err = fmt.Errorf("stee core has no attached store")
	}

	return c, err
}

func Store(v *viper.Viper) coreOption {
	storeType := v.GetString("type")
	return func(c *Core) error {
		newStore, err := storage.GetFactory(storeType)
		if err != nil {
			return err
		}
		c.store, err = newStore(v)
		return err
	}
}

// Close closes the core and the connection to its storage.
func (core *Core) Close() error {
	return core.store.Close()
}
