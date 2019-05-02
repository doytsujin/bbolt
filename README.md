# BBolt wrapper

This package define utilities functions for Bolt 

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

## Install package

```sh
go get github.com/lorenzodisidoro/bbolt
```

## Example

```golang
dir, _ := os.Getwd()
dbPath := dir + "/test.db"

// new instance
testBucket, err := bbolt.NewBBolt(dbPath, []byte("test_bucket"))
if err != nil {
    log.Fatal(err)
}

// put key/value into 'test_bucket'
testKey := []byte("love")
err = testBucket.Update(testKey, []byte("bitcoin"))
if err != nil {
    log.Fatal(err)
}

// get value by key
keyValue, err := testBucket.GetBy(testKey)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("I love %v", string(keyValue.Value))
```


