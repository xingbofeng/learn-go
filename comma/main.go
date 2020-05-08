package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("3243511.12341235"))
}

func comma(s string) string {
	var buf bytes.Buffer
	lastIndex := strings.LastIndex(s, ".")
	if lastIndex >= 0 {
		intPart := s[:lastIndex]
		insertIntPart(&buf, intPart)
		buf.WriteString(".")
		floatPart := s[lastIndex+1:]
		insertFloatPart(&buf, floatPart)
		return buf.String()
	}
	insertIntPart(&buf, s)
	return buf.String()
}

func insertIntPart(bufP *bytes.Buffer, intPart string) {
	length := len(intPart)
	for i, subString := range intPart {
		if (length-i)%3 == 0 && i != 0 {
			(*bufP).WriteString(",")
		}
		(*bufP).WriteString(string(subString))
	}
}

func insertFloatPart(bufP *bytes.Buffer, floatPart string) {
	for i, subString := range floatPart {
		if i%3 == 0 && i/3 != 0 {
			(*bufP).WriteString(",")
		}
		(*bufP).WriteString(string(subString))
	}
}
