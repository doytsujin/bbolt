package bbolt

import (
	"fmt"
	"os"
	"time"

	bolt "go.etcd.io/bbolt"
)

// BBolt define bbolt instance
type BBolt struct {
	Path       string
	BucketName []byte
	Storage    *bolt.DB
}

// KeyValue define response
type KeyValue struct {
	Key   []byte
	Value []byte
}

// NewBBolt function set up bbolt local storage
// @param path <string>: creates storage at the given path
// @param bucketName <[]byte>
// @return (BoltDatabase, error)
func NewBBolt(path string, bucketName []byte) (*BBolt, error) {
	bbolt := &BBolt{}
	options := &bolt.Options{Timeout: 10 * time.Second}
	var mode os.FileMode = 0600

	storage, err := bolt.Open(path, mode, options)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	bbolt.Path = path
	bbolt.Storage = storage
	bbolt.BucketName = bucketName

	return bbolt, nil
}

// Update create or update key parameter
// @param key <[]byte>
// @param value <[]byte>
// @return error
func (bb *BBolt) Update(key []byte, value []byte) error {
	err := bb.Storage.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bb.BucketName)
		if err != nil {
			return err
		}

		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

// GetBy retrieve data by key
// @param key <[]byte>
// @return error
func (bb *BBolt) GetBy(key []byte) (*KeyValue, error) {
	response := &KeyValue{}
	var value []byte
	var err error

	bb.Storage.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bb.BucketName)
		if bucket == nil {
			err = fmt.Errorf("Bucket %q not found", bb.BucketName)
			return err
		}

		value = bucket.Get(key)
		response.Key = key
		response.Value = value

		return nil
	})

	return response, err
}

// GetAll retrieves all bucket's key/value
// @return ([]string, error)
func (bb *BBolt) GetAll() ([]*KeyValue, error) {
	var response []*KeyValue

	bb.Storage.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bb.BucketName))
		if b != nil {
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				keyVal := &KeyValue{}
				keyVal.Key = k
				keyVal.Value = v

				response = append(response, keyVal)
			}
		}

		return nil
	})

	return response, nil
}
