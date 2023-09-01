package main

import (
	"fmt"
	"github.com/pkoukk/tiktoken-go"
	"github.com/tiktoken-go/tokenizer"
	"log"
)

func main1() {
	enc, err := tokenizer.Get(tokenizer.Cl100kBase)
	if err != nil {
		panic("oh oh")
	}

	// this should print a list of token ids
	ids, _, _ := enc.Encode("这是一个测试")
	fmt.Println(ids)

	// this should print the original string back
	text, _ := enc.Decode(ids)
	fmt.Println(text)
}

func main() {
	text := "这是一个测试"
	encoding := "cl100k_base"

	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return
	}

	// encode
	token := tke.Encode(text, nil, nil)

	//tokens
	log.Println(tke.Decode(token[:2]))
	// num_tokens
	log.Println(len(token))
}
