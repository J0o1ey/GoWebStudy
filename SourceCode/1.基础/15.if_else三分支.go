package main

import "fmt"

func main() {
	var a = 8
	if a < 8 {
		fmt.Println("虎六舅")
	} else if a == 8 {
		fmt.Println("老干妈") //else if可以无限添加
	} else { //else必须和右括号在同一行
		fmt.Println("往里整点哇哈哈")
	}
}
