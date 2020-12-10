package stee

import "github.com/milanrodriguez/stee/storage"

// Core is the central manager of Stee. Primary interaction point
type Core struct {
	store *storage.Store
}

// NewCore returns a new core.
func NewCore() *Core {
	storeParams := map[string]interface{}{"filepath": "./stee.db"}
	var store storage.Store
	store = &storage.IntegratedKV{}
	store.Open(storeParams)

	core := Core{}
	core.store = &store

	core.AddRedirection("_stee", "https://github.com/milanrodriguez/stee")

	return &core
}