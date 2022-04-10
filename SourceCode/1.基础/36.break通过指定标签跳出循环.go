package main

func main() {
OuterLoop: //外层循环标签
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 1:
				println(i, j)
				break OuterLoop
			case 2:
				println(i, j)
				break OuterLoop
			}
		}
	}
}
