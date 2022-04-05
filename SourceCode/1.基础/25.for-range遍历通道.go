package main

import "fmt"

func main() {
	c := make(chan int) //创建一个int型的通道
	go func() {         //启动一个goroutine
		c <- 7 //将数据推送进通道
		c <- 8
		c <- 9
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
