package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func processImage(width uint, height uint) image.Image {
	reader, err := os.Open("test.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	newJpeg := resize.Resize(width, height, m, resize.Lanczos3)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("rendered.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	err = jpeg.Encode(file, newJpeg, nil)
	if err != nil {
		log.Fatal(err)
	}

	return newJpeg
}

func main() {
	fmt.Println(processImage(500, 300))

	fmt.Println("uaas")
}
