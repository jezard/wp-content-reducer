package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type queueItem struct {
	filePath      string
	processedDate int
	thread        int
}

// test usage go run . /Users/jeremy/Library/
func main() {

	const env = "dev"
	var csvPath = ""

	if env == "dev" { //write to a directory where we can read it
		csvPath, _ = os.Getwd()
	} else {
		ex, _ := os.Executable()
		csvPath = filepath.Dir(ex)
	}

	if len(os.Args) < 2 {
		fmt.Println("Directory argument has not been provided")
		return
	}

	targetDir := os.Args[1]
	maxDepth := flag.Int("depth", 0, "Maximum recursion levels")

	err := os.Chdir(targetDir)

	if err != nil {
		fmt.Println("Directory argument is not a directory: ", targetDir)
	} else {
		// create a list or queue file
		cwd, _ := os.Getwd()

		f, err := os.Create(csvPath + "/queue.csv")
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

		// process queue
		processQueue(csvPath + "/queue.csv")

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
			m, _ := regexp.MatchString(".png|.jpeg|.jpg", strings.ToLower(e.Name()))
			if m {
				//filepath,status,thread
				fmt.Fprint(w, filepath.FromSlash("\""+dirName+"/"+e.Name())+"\",0,0\r\n")
				// fmt.Println(getIndent(depth), e.Name()) // logging
			}

		}
	}

	// directories
	for _, f := range dirEntries {

		if f.IsDir() {
			// fmt.Println(getIndent(depth), f.Name()+" (Dir)") // logging
			recurse(dirName+"/"+f.Name(), depth, maxDepth, w)
		}
	}
}

// process the queue of images
func processQueue(fileName string) {
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	fs := bufio.NewScanner(readFile)

	fs.Split(bufio.ScanLines)

	for fs.Scan() {

		fmt.Println(fs.Text())
	}

	readFile.Close()
}

// utility functions
func getIndent(depth int) string {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "-"
	}
	return indent
}
