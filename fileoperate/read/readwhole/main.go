package main

import (
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("文件内容：", string(content))
}
