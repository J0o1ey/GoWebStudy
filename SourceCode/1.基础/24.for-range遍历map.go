package main

import (
	"fmt"
)

func main() {
	m := map[string]int{ //键为string，键值为int
		"go":  100,
		"web": 100,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
}
