package main

import (
	"bgskoro21/be-pos/tools/generator"
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <name>")
		return
	}

	name := os.Args[1]
	generator.GenerateRepository(name)
}