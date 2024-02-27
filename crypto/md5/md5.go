package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
