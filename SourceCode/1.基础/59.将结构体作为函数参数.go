package main

import "fmt"

type Cyberspace_mapping_engine struct {
	name    string
	company string
	score   int
}

func print_engine(engine Cyberspace_mapping_engine) { //创建输出函数
	fmt.Printf("网络空间测绘引擎公司为%s,评分为%d \n", engine.company, engine.score)
}
func main() {
	var fofa Cyberspace_mapping_engine
	var Hunter Cyberspace_mapping_engine

	//描述fofa
	fofa.company = "华顺信安"
	fofa.score = 100

	//描述Hunter
	Hunter.company = "奇安信"
	Hunter.score = 100

	//调用输出函数，将结构体作为参数，打印fofa和Hunter的信息
	print_engine(fofa)
	print_engine(Hunter)
}
