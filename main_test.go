package bbolt

import (
	"log"
	"os"
	"testing"
)

const (
	k = "k1"
	v = "v1"
)

func TestEncryptAndDecrypt(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := dir + "/test.db"
	testBucket, err := NewBBolt(dbPath, []byte("test_bucket"))
	if err != nil {
		log.Fatal(err)
	}

	err = testBucket.Update([]byte(k), []byte(v))
	if err != nil {
		log.Fatal(err)
	}

	keyValue, err := testBucket.GetBy([]byte(k))
	if err != nil {
		log.Fatal(err)
	}

	if string(keyValue.Key) != k || string(keyValue.Value) != v {
		log.Fatal("Should be expected value = " + v)
	}
}
