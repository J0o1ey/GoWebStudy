package main

import (
	"fmt"
	"strconv"
)

var (
	age  int
	name string
)

func main() {
	name = "TOM"
	age = 11
	age := strconv.Itoa(age) //把age转换成string型
	fmt.Println(name + " is " + age + " years old")
}
