package main

import "fmt"

func main() {
	for key, value := range []int{0, 1, -1, -2} {
		fmt.Printf("key:%d value:%d\n", key, value)
	}
}
