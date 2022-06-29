package main

import "fmt"

func main() {
	var arr [6]int //声明一个长度为6的数组
	var i, j int
	for i = 0; i < 6; i++ {
		arr[i] = i + 66 //设置元素为i+66
	}
	for j = 0; j < 6; j++ { //输出每个数组元素的值
		fmt.Printf("Array[%d] = %d\n", j, arr[j])
	}
}
