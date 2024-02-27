package ace

import (
	"bytes"
	"testing"

	"github.com/lankaiyun/kaiyunchain/crypto/md5"
)

func TestACE(t *testing.T) {
	content := "Secret Content"
	testData := []struct{ passwd string }{
		{"1"},
		{"11112222"},
		{"aaaa1111"},
		{"abcdefgf"},
		{"./#%(@#*#da"},
	}
	for _, tt := range testData {
		contentB := []byte(content)
		keyB := []byte(md5.Md5(tt.passwd))
		cipherText, _ := AESEncrypt(contentB, keyB)
		actual, _ := AESDecrypt(cipherText, keyB)
		if !bytes.Equal(actual, contentB) {
			t.Errorf("Decryption Failure")
		}
	}
}
