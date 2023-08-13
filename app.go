package main

import "fmt"

func main() {
	a := 1
	var b string = "hello"

	fmt.Println("hello", a)
	var c = fmt.Sprintf("my name is :%s AND vikki %d", b, a)
	fmt.Println(c)
}
