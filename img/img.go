package img

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"log"
	"os"
)

func ProcessImage(width uint, height uint) *bytes.Buffer {
	reader, err := os.Open("./assets/unicorn2.png")
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

	//	file, err := os.Create("rendered.jpeg")
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	buff := new(bytes.Buffer)

	err = png.Encode(buff, newJpeg)
	if err != nil {
		log.Fatal(err)
	}

	//	err = jpeg.Encode(file, newJpeg, nil)
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	return buff
}
