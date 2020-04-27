package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	// 通过buf操作字符串
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		// 通过Fprintf将数字写入buf底层
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
