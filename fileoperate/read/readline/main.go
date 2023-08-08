package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		log.Println("读取到的内容：", scanner.Text())
	}
}
