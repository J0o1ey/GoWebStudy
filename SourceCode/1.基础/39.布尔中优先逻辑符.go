package main

import "fmt"

func main() {
	var c int = 100
	if 1 <= c && c <= 9 || 10 <= c && c <= 90 || 100 <= c && c <= 900 {
		fmt.Println("c在范围内")
	}
}
