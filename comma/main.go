package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1112345"))
}

func comma(s string) string {
	var buf bytes.Buffer
	length := len(s)
	for i, subString := range s {
		if (length-i)%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteString(string(subString))
	}
	return buf.String()
}
