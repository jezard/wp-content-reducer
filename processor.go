package main

import (
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

func processImage(imagePath string) bool {
	reader, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	if strings.HasSuffix(imagePath, ".jpeg") || strings.HasSuffix(imagePath, ".jpeg") {
		m, err := jpeg.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(m.Bounds().String())

	} else if strings.HasSuffix(imagePath, ".png") {
		m, err := png.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(m.Bounds().String())
	} else if strings.HasSuffix(imagePath, ".gif") {
		m, err := gif.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(m.Bounds().String())
	}
	fmt.Println("")

	return false
}
