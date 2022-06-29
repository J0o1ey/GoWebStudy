package main

import "fmt"

func main() {
	str := "Golang Web"
	for i := 1; i < 5; i++ {
		str += "Security"
		fmt.Println(str)
	}
}
