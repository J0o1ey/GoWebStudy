package main

import "fmt"

func main() {
	str := "I love Golang Security Programming"
	chars := []rune(str) //把字符串转化为rune切片
	for _, char := range chars {
		fmt.Println(string(char))
	}
}
