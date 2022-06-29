package main

import (
	"fmt"
)

func main() {
	str := "Golang Security Programming"
	for index, char := range str {
		fmt.Println("%d %c \n", index, char) //输出索引、以单个字符的ascii码
	}
}
