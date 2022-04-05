package main

import "fmt"

func main() {
	m := map[string]int{
		"fofa":   100,
		"shodan": 90,
	}
	for _, value := range m {
		fmt.Println(value)
	}
}
