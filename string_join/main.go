package main

import (
	"bytes"
	"fmt"
	"strings"
)

const base = "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

var baseSlice []string

func init() {
	for i := 0; i < 200; i++ {
		baseSlice = append(baseSlice, base)
	}
}

func main() {

}

func SingleSumString(str string) string {
	return base + str
}

func SingleSprintfString(str string) string {
	return fmt.Sprintf("%s%s", base, str)
}

func SingleBuilderString(str string) string {
	var builder strings.Builder
	builder.Grow(2 * len(str))
	builder.WriteString(base)
	builder.WriteString(str)
	return builder.String()
}

func SingleBytesString(str string) string {
	buf := new(bytes.Buffer)
	buf.WriteString(base)
	buf.WriteString(str)
	return buf.String()
}

func SingleByteSliceString(str string) string {
	buf := make([]byte, 0)
	buf = append(buf, base...)
	buf = append(buf, str...)
	return string(buf)
}

func SingleJoinstring() string {
	return strings.Join([]string{base, base}, "")
}

func SumString() string {
	res := ""
	for _, val := range baseSlice {
		res += val
	}
	return res
}

func SprintfString() string {
	res := ""
	for _, val := range baseSlice {
		res = fmt.Sprintf("%s%s", res, val)
	}
	return res
}

func BuilderString() string {
	var builder strings.Builder
	builder.Grow(200 * len(baseSlice))
	for _, val := range baseSlice {
		builder.WriteString(val)
	}
	return builder.String()
}

func bytesString() string {
	buf := new(bytes.Buffer)
	for _, val := range baseSlice {
		buf.WriteString(val)
	}
	return buf.String()
}

func byteSliceString() string {
	buf := make([]byte, 0)
	for _, val := range baseSlice {
		buf = append(buf, val...)
	}
	return string(buf)
}

func Joinstring() string {
	return strings.Join(baseSlice, "")
}

func main4() {
	buf := make([]byte, 0)
	base := "Hello world"
	buf = append(buf, base...)
	fmt.Println(string(buf))
}

func main3() {
	slice := []string{"Hello", "World"}
	fmt.Println(strings.Join(slice, ""))
}

func main2() {
	buf := new(bytes.Buffer)
	buf.WriteString("Hello World")
	fmt.Println(buf.String())
}

func main1() {
	var builder strings.Builder
	builder.WriteString("Hello world")
	fmt.Println(builder.String())
}
