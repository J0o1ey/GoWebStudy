package main

import "fmt"

type Cyberspace_mapping_engine struct {
	name    string
	comapny string
	score   int
}

func main() {
	var fofa Cyberspace_mapping_engine   //声明fofa为Cyberspace_mapping_engine类型
	var Hunter Cyberspace_mapping_engine //声明Hunter为Cyberspace_mapping_engine类型

	//fofa描述
	fofa.name = "fofa"
	fofa.comapny = "华顺信安"
	fofa.score = 100

	//Hunter描述
	Hunter.name = "Hunter"
	Hunter.comapny = "奇安信"
	Hunter.score = 100

	//打印上述信息
	fmt.Printf("fofa_conpamy: %s fofa_score: %d \n", fofa.comapny, fofa.score)
	fmt.Printf("Hunter_company %s hunter_score: %d \n", Hunter.comapny, Hunter.score)
}
