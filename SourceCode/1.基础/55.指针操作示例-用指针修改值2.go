package main

import "fmt"

func exchange(a, b *int) {
	a, b = b, a
}
func main() {
	x, y := 6, 9
	exchange(&x, &y)
	fmt.Println(x, y)
}
