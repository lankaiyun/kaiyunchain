package common

import (
	"bufio"
	"log"
	"os"
)

func WriteFile(filename, content string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panic(err)
		}
	}(file)
	if err != nil {
		log.Panic("Failed to OpenFile: ", err)
	}
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		log.Panic(err)
	}
	err = writer.Flush()
	if err != nil {
		log.Panic(err)
	}
}

func ReadFile(filename string) []byte {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Panic("Failed to ReadFile: ", err)
	}
	return content
}
