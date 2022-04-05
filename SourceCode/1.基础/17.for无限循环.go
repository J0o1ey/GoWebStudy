package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		fmt.Println(i)
		if i > 50 {
			break
		}
	}
}
