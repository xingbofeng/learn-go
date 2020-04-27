package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// strings.Join就是join方法，第一个参数传数据，第二个参数传连接字符串
	fmt.Println(strings.Join(os.Args[1:], " "))
}
