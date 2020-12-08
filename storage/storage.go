package storage

type Driver interface {
	Open(map[string]interface{}) (error)
	Close() (error)
	ReadRedirection(string) (string, bool)
	WriteRedirection(string, string) (error)
	DeleteRedirection(string) (error)
}