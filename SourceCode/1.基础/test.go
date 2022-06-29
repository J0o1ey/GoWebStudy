package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for x := 0; x < 10; x++ {
			println(i, x)
			if i == 2 {
				println("done")
				goto breakTag
			}
		}
	}
	return
breakTag:
	fmt.Println("执行完成")
}
