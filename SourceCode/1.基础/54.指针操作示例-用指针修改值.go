package main

import "fmt"

func exchange(c, d *int) {
	t := *c //取c的指针，赋给临时变量t
	*c = *d //将d指针的值，赋给c指针指向的变量
	*d = t  //将临时变量t的值赋给d指针指向的变量
}
func main() {
	a, b := 6, 9      //准备两个变量，赋值6，9
	exchange(&a, &b)  //交换变量值
	fmt.Println(a, b) //输出变量值
}
