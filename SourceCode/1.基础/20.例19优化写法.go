package main

func main() {
	var i int
JumpLoop:
	for {
		if i >= 10 {
			break JumpLoop
		}
		i++
		println(i)
	}

}
