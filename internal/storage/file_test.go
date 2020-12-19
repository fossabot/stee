package storage

import (
	"bytes"
	"testing"

	"github.com/spf13/viper"
)

func TestIntegratedKV(t *testing.T) {

	var err error
	var key string = t.Name()
	var target string = t.Name() + "//redirection"

	var config = []byte(`
type: "file"
path: ` + t.TempDir() + "/stee.db")

	v := viper.New()
	v.ReadConfig(bytes.NewBuffer(config))

	store, err := newFileStore(v)
	if err != nil {
		t.Errorf("Could not create the store")
	}

	// open

	err = store.Open()
	if err != nil {
		t.Errorf("Could not open the store")
	}

	// write

	err = store.WriteRedirection(key, target)
	if err != nil {
		t.Errorf("Failed to write in the store")
	}

	// read

	readValue, err := store.ReadRedirection(key)
	if err != nil {
		t.Errorf("Failed to read in the store")
	}
	if readValue != target {
		t.Errorf("Wrong value from the store")
	}

	// delete

	err = store.DeleteRedirection(key)
	if err != nil {
		t.Errorf("Failed to delete in the store")
	}

	// read (should fail)

	_, err = store.ReadRedirection(key)
	if err == nil {
		t.Errorf("Succeed to read value after delete. Shouldn't have")
	}

	// close

	store.Close()

	// Try to read again

	_, err = store.ReadRedirection(key)
	if err == nil {
		t.Errorf("Succeed to read value after delete. Shouldn't have")
	}
}
