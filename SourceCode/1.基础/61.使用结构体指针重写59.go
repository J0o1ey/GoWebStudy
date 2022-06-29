package main

import "fmt"

type Cyberspace_mapping_engine struct {
	name    string
	company string
	score   int
}

func print_engine(engine *Cyberspace_mapping_engine) {
	fmt.Printf("网络空间测绘引擎名称%s,开发公司%s,得分%d\n", engine.name, engine.company, engine.score)
	fmt.Printf("结构体的内存地址为%p\n", &engine)
}
func main() {
	var fofa Cyberspace_mapping_engine
	var Hunter Cyberspace_mapping_engine

	//fofa描述
	fofa.name = "fofa"
	fofa.company = "华顺信安"
	fofa.score = 100

	//Hunter描述
	Hunter.name = "Hunter"
	Hunter.company = "奇安信"
	Hunter.score = 100

	print_engine(&fofa)
	print_engine(&Hunter)
}