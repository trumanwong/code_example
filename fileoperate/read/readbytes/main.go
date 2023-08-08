package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	for {
		n, err := r.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		log.Println("读取到的字节数：", n)
		// 分字节读取文件时，最后打印的时候应该调用string(buffer[:n])而不是调用string(buffer)来获取读取到的值，否则可能会打印出额外的字节。
		log.Println("读取到的内容：", string(buffer[:n]))
	}
}
