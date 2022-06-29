package main

import "fmt"

type Cyberspace_mapping_engine struct {
	name    string
	company string
	score   int
}

func main() {
	structPointer := &Cyberspace_mapping_engine{"fofa", "华顺信安", 100}
	fmt.Printf("结构体变量内存地址%p \n", structPointer)
}
