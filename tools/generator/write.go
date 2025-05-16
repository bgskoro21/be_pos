package generator

import (
	"fmt"
	"os"
)

func write(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created:", path)
}
