package storage

import (
	"errors"
	"fmt"

	bolt "go.etcd.io/bbolt"
)

const bucketName string = "root"

var ErrMissingParamFilepath error = errors.New("Could not find parameter \"filepath\"")

// IntegratedKV is an embedded KV store. It stores data in a file.
type IntegratedKV struct {
	db *bolt.DB
}

func (i *IntegratedKV) Open(params map[string]interface{}) error {
	path, ok := params["filepath"].(string)
	if !ok || path == "" {
		panic(ErrMissingParamFilepath)
	}

	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return err
	}

	i.db = db

	i.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return err
}

func (i *IntegratedKV) Close() error {
	return i.db.Close()
}

func (i *IntegratedKV) ReadRedirection(key string) (string, bool) {
	var result []byte
	i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		result = b.Get([]byte(key))
		return nil
	})
	ok := true
	if result == nil {
		ok = false
	}
	return string(result), ok
}

func (i *IntegratedKV) WriteRedirection(key string, target string) (err error) {
	i.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Put([]byte(key), []byte(target))
		return err
	})
	return err
}

func (i *IntegratedKV) DeleteRedirection(key string) (err error) {
	i.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Delete([]byte(key))
		return err
	})
	return err
}
