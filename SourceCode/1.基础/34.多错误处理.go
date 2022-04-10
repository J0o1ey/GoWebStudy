package main

import "fmt"

func main() {
	//......省略前面代码
	err := getUserInfo()
	if err != nil {
		fmt.Println(err)
		exitProcess()
		return
	}
	err = getEmail()
	if err != nil {
		fmt.Println(err)
		exitProcess()
		return
	}
	fmt.Println("over")

}
