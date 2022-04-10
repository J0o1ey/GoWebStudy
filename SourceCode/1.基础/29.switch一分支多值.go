package main

import "fmt"

func main() {
	var language = "python"
	switch language {
	case "golang", "java", "python":
		fmt.Println("I love", language)
	}
}
