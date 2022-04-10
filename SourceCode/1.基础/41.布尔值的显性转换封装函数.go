package main

import "fmt"

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	result := BoolToInt(false)
	fmt.Println(result)
	result2 := BoolToInt(true)
	fmt.Println(result2)
}
