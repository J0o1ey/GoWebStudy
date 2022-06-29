package main

import "fmt"

func main() {
	var address = "Xian China"
	ptr := &address
	fmt.Printf("ptr type: %T \n", ptr)
	fmt.Printf("address %p \n", ptr)

	value := *ptr //对指针进行取值操作
	fmt.Printf("value type %T \n", value)
	fmt.Printf("value address %s \n", value)
}
