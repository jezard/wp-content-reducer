package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	recurse("/Users/jeremy/Desktop/CE Project", -1)
}

func recurse(dirName string, depth int) {
	depth++
	if depth > 4 {
		return
	}

	dirEntries, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error: " + err.Error() + " Dir: " + dirName + " at depth " + strconv.Itoa(depth))
	}

	//files
	for _, e := range dirEntries {
		if !e.IsDir() {
			fmt.Println(getIndent(depth), e.Name())
		}
	}

	// directories
	for _, f := range dirEntries {

		if f.IsDir() {
			fmt.Println(getIndent(depth), f.Name()+" (Dir)")
			recurse(dirName+"/"+f.Name(), depth)
		}
	}
}

func getIndent(depth int) string {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "-"
	}
	return indent
}
