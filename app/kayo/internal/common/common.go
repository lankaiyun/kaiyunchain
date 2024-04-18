package common

import (
	"os"
	"strconv"
	"unicode"
)

func IsInitDir() bool {
	_, err := os.Stat("./KaiYunChainData")
	return !os.IsNotExist(err)
}

func IsContainsDigitAndLetter(s string) bool {
	hasDigit := false
	hasLetter := false

	for _, c := range s {
		if unicode.IsDigit(c) {
			hasDigit = true
		}
		if unicode.IsLetter(c) {
			hasLetter = true
		}
		if hasDigit && hasLetter {
			return true
		}
	}

	return false
}

func IsPositiveInteger(s string) bool {
	num, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return num > 0
}

func IsInteger(s string) bool {
	// 尝试将字符串转换为整数
	_, err := strconv.Atoi(s)
	if err != nil {
		// 转换失败，说明不是整数
		return false
	}
	// 判断是否为正整数
	return true
}
