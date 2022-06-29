package main

import "fmt"

func main() {
	str := "I love Golang安全开发"
	by := []byte(str) //转换为[]byte，数据将自动复制
	by[2] = ' '       //把第二个字节转化为空格
	fmt.Printf("%s \n", str)
	fmt.Printf("%s\n", by)
}
