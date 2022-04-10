package main

import (
	"fmt"
)

func main() {
	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			if y == 2 {
				goto breakTag //跳转到breakTag标签
			}
		}
	}
	return
breakTag: //定义breakTag标签
	fmt.Println("done")
}
