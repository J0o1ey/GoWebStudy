package main

import (
	"bytes"
	"fmt"
)

func getNextString() string {
	//here is your code
}

func main() {
	var buffer bytes.Buffer
	for {
		if piece, ok := getNextString(); ok {
			//通过WriteString方法，将需要串联的字符串写入buffer
			buffer.WriteString(piece)
		} else {
			break
		}
	}
	fmt.Println(buffer.String()) //取回整个级联的字符串
}
