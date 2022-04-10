package main

import "fmt"

func main() {
	var r int = 6
	switch {
	case r > 1 && r < 10: //多分支表达式判断
		fmt.Println(r)
	case r < 20:
		fmt.Println("r可以使用") //执行到上一句判断成功就停止，本句不会执行
	}
}
