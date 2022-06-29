package main

import "fmt"

func main() {
	str := "hello 安全"
	ru := []rune(str)
	ru[6] = '中' //不可以用双引号
	ru[7] = '国'
	fmt.Println(str)
	fmt.Println(string(ru))
}
