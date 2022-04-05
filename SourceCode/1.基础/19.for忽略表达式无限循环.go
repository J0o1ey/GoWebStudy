package main

func main() {
	var i int
JumpLoop:
	for ; ; i++ { //for没有设置i的初始值，也没有设置i的循环条件表达式，此时循环会一直持续下去
		println(i)
		if i >= 10 {
			break JumpLoop // i大于10的时候通过break跳出JumpLoop标签对应的for循环
		}
	}
}
