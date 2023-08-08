package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	n, err := f.WriteString("Hello, world!")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("写入的字节数：", n)
}
