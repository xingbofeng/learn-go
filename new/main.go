package main

import "fmt"

func main() {
	var temp int
	p := &temp
	fmt.Println(*p)
	*p++
	fmt.Println(*p)
	fmt.Println(temp)
}
