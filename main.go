package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
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

		f, err := os.Create("queue.csv")
		if err != nil {
			fmt.Println("Error: Could not open queue file.")
		}

		defer f.Close()

		w := bufio.NewWriter(f)
		fmt.Fprint(w, "Filepath,Status,Thread\r\n")

		// let's begin!
		recurse(cwd, 0, *maxDepth, w)

		err = w.Flush()

		if err != nil {
			fmt.Println("Couldn't write queue from buffer to file: ", err)
		}
	} else {
		fmt.Println("Directory argument is not a directory")
	}
}

func recurse(dirName string, depth int, maxDepth int, w *bufio.Writer) {
	depth++

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
			//filepath,status,thread
			fmt.Fprint(w, filepath.FromSlash(dirName+"/"+e.Name())+",0,0\r\n")
			fmt.Println(getIndent(depth), e.Name())
		}
	}

	// directories
	for _, f := range dirEntries {

		if f.IsDir() {
			fmt.Println(getIndent(depth), f.Name()+" (Dir)")
			recurse(dirName+"/"+f.Name(), depth, maxDepth, w)
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
