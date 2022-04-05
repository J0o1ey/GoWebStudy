package main

import "fmt"

func main() {
	var str = "hi 哈哈嗨"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}
}
