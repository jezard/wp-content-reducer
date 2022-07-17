package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type queueItem struct {
	filePath    string
	isProcessed string
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
		// if the queue file does not exist, create it.
		if _, err := os.Stat(csvPath + "/queue.csv"); errors.Is(err, os.ErrNotExist) {
			// create a list or queue file
			cwd, _ := os.Getwd()

			f, err := os.Create(csvPath + "/queue.csv")
			if err != nil {
				fmt.Println("Error: Could not open queue file.")
			}

			defer f.Close()

			w := bufio.NewWriter(f)

			// let's begin!
			walkDir(cwd, w)

			err = w.Flush()

			if err != nil {
				fmt.Println("Couldn't write queue from buffer to file: ", err)
			}
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
				fmt.Fprint(w, path+"|0\r\n")
			}

			if err != nil {
				return err
			}
			return nil
		})
	return err
}

// process the queue of images
func processQueue(fileName string) {

	queueFile, err := os.OpenFile(fileName, os.O_RDWR, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
	}
	fs := bufio.NewScanner(queueFile)

	fs.Split(bufio.ScanLines)

	pos := -1 // start position

	for fs.Scan() {
		len := len(fs.Bytes())
		pos += len
		sa := strings.Split(fs.Text(), "|")
		item := queueItem{sa[0], sa[1]}

		// if not yet processed, process image
		if item.isProcessed == "0" {

			// processImage
			if processImage(item.filePath) { // successfully processed, mark the entry done.
				queueFile.WriteAt([]byte("1"), int64(pos))
			}
		}
		pos += 2 // to account for the \r\n ?
	}
	queueFile.Close()
}
