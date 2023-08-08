package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	arr := []string{
		"Hello World",
		"Hello Golang",
	}

	for _, line := range arr {
		n, err := fmt.Fprintln(f, line)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("写入的字节数：", n)
		log.Println("写入的内容：", line)
	}
	log.Println("写入完成")
}
