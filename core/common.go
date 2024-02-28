package core

import (
	"bytes"
	"encoding/gob"
)

func Serialize(i interface{}) []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(i)
	return buffer.Bytes()
}
