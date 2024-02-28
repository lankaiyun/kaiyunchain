package log

import (
	"log"
	"os"
)

func GetLogger(path string) *log.Logger {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	return log.New(f, "", log.LstdFlags|log.Llongfile)
}
