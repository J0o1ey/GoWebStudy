package main

import (
	"fmt"
)

func main() {
	var isBreak bool
	for x := 0; x < 20; x++ { //外循环
		for y := 0; y < 20; y++ { //内循环
			if y == 2 { //设置某个条件时退出循环
				isBreak = true //设置退出标记
				break          //退出本次循环
			}
		}
		if isBreak {
			break //根据标记，还需要推出一次循环
		}
	}
	fmt.Println("over")
}
