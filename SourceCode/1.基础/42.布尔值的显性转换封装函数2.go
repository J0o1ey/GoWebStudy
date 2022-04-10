package main

import "fmt"

func IntToBool(i int) bool {
	return i != 0
}
func main() {
	result := IntToBool(1)
	fmt.Println(result)

	result2 := IntToBool(0)
	fmt.Println(result2)
}
