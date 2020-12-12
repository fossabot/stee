package storage

import (
	"testing"
)

func TestIntegratedKV(t *testing.T) {

	var err error
	var key string = t.Name()
	var target string = t.Name() + "//redirection"

	store := IntegratedKV{}

	storeParams := map[string]interface{}{"filepath": t.TempDir() + "/stee.db"}

	// open

	err = store.Open(storeParams)
	if err != nil || store.db == nil {
		t.Errorf("Could not open the store")
	}

	// write

	err = store.WriteRedirection(key, target)
	if err != nil {
		t.Errorf("Failed to write in the store")
	}

	// read

	readValue, ok := store.ReadRedirection(key)
	if ok != true {
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

	_, ok = store.ReadRedirection(key)
	if ok != false {
		t.Errorf("Succeed to read value after delete. Shouldn't have")
	}

	// close

	store.Close()

	// Try to read again

	_, ok = store.ReadRedirection(key)
	if ok != false {
		t.Errorf("Succeed to read value after delete. Shouldn't have")
	}

	// Try to open without params

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("No panic when not providing params to open store")
		}
	}()
	err = store.Open(map[string]interface{}{})
}
