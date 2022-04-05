package main

import (
	"fmt"
)

func main() {
jumpLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			if i > 2 {
				break jumpLoop
			}
			fmt.Println(i)
		}
	}
}
