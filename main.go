package main

import (
	"fmt"
	"os"
)

func main() {
	dirs, _ := os.ReadDir(".")
	for i, e := range dirs {
		fmt.Println(i, e.Name())
	}
}
