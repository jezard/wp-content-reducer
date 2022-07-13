package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
		fmt.Fprint(w, "Filepath|Status|Thread\r\n")

		// let's begin!
		walkDir(cwd, w)

		err = w.Flush()

		if err != nil {
			fmt.Println("Couldn't write queue from buffer to file: ", err)
		}

		// process queue
		processQueue(csvPath + "/queue.csv")

	}
}

func walkDir(cwd string, w *bufio.Writer) error {
	err := filepath.Walk(cwd,
		func(path string, info os.FileInfo, err error) error {

			m, _ := regexp.MatchString(".png|.jpeg|.jpg", strings.ToLower(info.Name()))
			if m {
				//filepath,status,thread
				fmt.Fprint(w, path+"|0|0\r\n")
			}

			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	return err
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

		//fmt.Println(fs.Text())
	}

	readFile.Close()
}
