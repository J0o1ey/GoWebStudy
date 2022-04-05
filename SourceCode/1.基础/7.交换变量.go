package main

import "fmt"

func main() {
	a, b := "a", "b"
	b, a = a, b
	fmt.Println(a, b)
}
