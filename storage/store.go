package storage

// Store is an interface representing the possible interactions with a storage.
// When working with a Store, you should always open it before use and close it after use.
type Store interface {
	// Open initialize the store. (Could be opening a file or initiating a connection to a remote database)
	Open(map[string]interface{}) error
	// Close closes the store. (Could be closing a file or closing a remote connection.)
	Close() error
	// ReadRedirection takes a redirection key and returns the redirection's target and a success boolean.
	// If it fails, it returns "" as target.
	ReadRedirection(string) (string, bool)
	// WriteRedirection writes a redirection to the store. Parameters are the key and the target.
	WriteRedirection(string, string) error
	// DeleteRedirection deletes a redirection from the store. it Takes the key as argument.
	DeleteRedirection(string) error
}
