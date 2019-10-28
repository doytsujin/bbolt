# BBolt wrapper

[![Build Status](https://travis-ci.com/lorenzodisidoro/bbolt.svg?branch=master)](https://travis-ci.com/lorenzodisidoro/bbolt)

This package define utilities functions for Bolt.

## Install package

```sh
go get github.com/lorenzodisidoro/bbolt
```

## Usage
Import `bbolt` package in your GO file
```golang
import "github.com/lorenzodisidoro/bbolt"
```

it define `BBolt` and `KeyValue` types
```golang
// BBolt define instance
type BBolt struct {
	Path       string
	BucketName []byte
	Storage    *bolt.DB
}

// KeyValue define functions response
type KeyValue struct {
	Key   []byte
	Value []byte
}
```

### New key-value store
```golang
dir, _ := os.Getwd()
dbPath := dir + "/test.db"

// new BBolt instance
testBucket, err := bbolt.NewBBolt(dbPath, []byte("test_bucket"))
if err != nil {
    log.Fatal(err)
}
```

### Add new key-value or update
```golang
// put key/value into 'test_bucket'
testKey := []byte("love")
err = testBucket.Update(testKey, []byte("bitcoin"))
if err != nil {
    log.Fatal(err)
}
```

### Get by key
```golang
// get value by key
keyValue, err := testBucket.GetBy(testKey)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("I love %v", string(keyValue.Value))
```
