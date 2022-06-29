package main

import "fmt"

func main() {
	var score = 100
	var name string = "Tom"
	//用fmt.Println函数的动词"%p"，输出对score和name变量取地址后的指针值
	fmt.Printf("%p %p \n", &score, &name)
}
