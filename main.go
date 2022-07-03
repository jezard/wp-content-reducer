package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Directory argument has not been provided")
		return
	}

	cwd, _ := os.Getwd()
	targetDir := os.Args[1]
	maxDepth := flag.Int("depth", 0, "Maximum recursion levels")

	err := os.Chdir(cwd + "/" + targetDir)

	if err != nil {
		cwd, _ := os.Getwd()
		recurse(cwd, 0, *maxDepth)
	} else {
		fmt.Println("Directory argument is not a directory")
	}
}

func recurse(dirName string, depth int, maxDepth int) {
	depth++
	//fmt.Println(depth, maxDepth)
	if maxDepth > 0 && depth > maxDepth {
		return
	}

	dirEntries, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error: " + err.Error() + " Dir: " + dirName + " at depth " + strconv.Itoa(depth))
	}

	// files
	for _, e := range dirEntries {
		if !e.IsDir() {
			fmt.Println(getIndent(depth), e.Name())
		}
	}

	// directories
	for _, f := range dirEntries {

		if f.IsDir() {
			fmt.Println(getIndent(depth), f.Name()+" (Dir)")
			recurse(dirName+"/"+f.Name(), depth, maxDepth)
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
