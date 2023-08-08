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
	content := []byte("Hello, world!")
	i := 0
	byteLen := 1024
	for {
		if i*byteLen > len(content) {
			break
		}

		start, end := i*byteLen, (i+1)*byteLen
		if (i+1)*byteLen > len(content) {
			start, end = i*byteLen, len(content)
		}
		n, err := f.Write(content[start:end])
		if err != nil {
			log.Fatal(err)
		}
		i++
		log.Println("写入的字节数：", n)
	}
	log.Println("写入完成")
}
