package keccak256

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func ExampleKeccak256() {
	b := Keccak256([]byte("hello world"))
	fmt.Printf("十六进制字符串编码：%s\n", hex.EncodeToString(b))
	fmt.Printf("十六进制字符串长度：%d\n", len(hex.EncodeToString(b)))
	fmt.Printf("byte长度：%d\n", len(b))
	// Output:
	// 十六进制字符串编码：47173285a8d7341e5e972fc677286384f802f8ef42a5ec5f03bbfa254cb01fad
	// 十六进制字符串长度：64
	// byte长度：32
}

func TestName(t *testing.T) {

}
