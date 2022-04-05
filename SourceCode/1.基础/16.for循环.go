package main

import "fmt"

func main() {
	id := 1
	for i := 1; i < 5; i++ {
		id *= i
		fmt.Println(id)
	}
}
