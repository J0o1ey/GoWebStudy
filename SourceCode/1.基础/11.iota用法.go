package main

import "fmt"

type Direction int

const (
	North Direction = iota
	East
	West
	South
)

func main() {
	fmt.Println(East)
}
