package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}
		// 通过ioutil读取的文件是一个多行字符串，通过\n来进行分割
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		fmt.Println(n, line)
	}
}
