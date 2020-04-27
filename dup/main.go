package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 通过make可以创建map，key是string类型，value是int类型
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// 即使最开始没有值也会被初始化为int的默认值0
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}
}
