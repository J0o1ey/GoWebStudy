package main

import "fmt"

func main() {
	//声明三个局部变量
	var localvar1, localvar2, localvar3 int

	//变量赋值
	localvar1 = 2
	localvar2 = 3
	localvar3 = localvar1 + localvar2
	
	fmt.Printf("localvar1 = %d,localvar2 = %d,localvar3 = %d\n", localvar1, localvar2, localvar3)
}
