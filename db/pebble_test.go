package db

import (
	"fmt"
)

func Example() {
	db := GetDbObj("./testdata")
	Set([]byte("hello"), []byte("world"), db)
	fmt.Println(string(Get([]byte("hello"), db)))
	// Output:
	// world
	CloseDbObj(db)
}
