package main

import "fmt"

type Cyberspace_mapping_engine struct {
	name    string
	company string
	score   int
} //声明结构体

func main() {
	//创建一个新的结构体
	fmt.Println(Cyberspace_mapping_engine{"fofa", "华顺信安", 100})

	//使用key => value模式
	fmt.Println(Cyberspace_mapping_engine{name: "Hunter", company: "奇安信", score: 100})

	//忽略的字段为0或空
	fmt.Println(Cyberspace_mapping_engine{name: "Quake", company: "360"})
}
