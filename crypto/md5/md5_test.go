package md5

import (
	"encoding/hex"
	"fmt"
)

func ExampleMd5() {
	str := Md5("hello world")
	b, _ := hex.DecodeString(str)
	fmt.Printf("十六进制字符串编码：%s\n", str)
	fmt.Printf("十六进制字符串长度：%d\n", len(str))
	fmt.Printf("byte长度：%d\n", len(b))
	// Output:
	// 十六进制字符串编码：5eb63bbbe01eeed093cb22bb8f5acdc3
	// 十六进制字符串长度：32
	// byte长度：16
}
