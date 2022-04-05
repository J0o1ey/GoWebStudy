package main

import "fmt"

func main() {
	for key, _ := range []int{9, 8, 7, 6} {
		fmt.Printf("key:%d \n ", key)
	}
}
