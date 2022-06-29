package main

import "fmt"

func main() {
	var SliceBuilder [20]int
	for i := 0; i < 20; i++ {
		SliceBuilder[i] = i + 1
	}
	fmt.Println(SliceBuilder[0:5]) //区间元素
	fmt.Println(SliceBuilder[15:]) //从中间到尾部的所有元素
	fmt.Println(SliceBuilder[:5])  //从开头到指定位置的所有元素
}
