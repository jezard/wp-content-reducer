package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	dirs, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i, e := range dirs {
		fmt.Println(i, e.Name())
		if e.IsDir() { // TODO should be recursive
			os.Chdir(e.Name())
			dirs1, err := os.ReadDir(".")

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			for j, f := range dirs1 {
				fmt.Println(strconv.Itoa(i)+"."+strconv.Itoa(j), "_"+f.Name())
			}
		}
	}
}
