package main

import (
	"image"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
)

func main() {
	// 100 x 200 pixels
	myImage := image.NewRGBA(image.Rect(0, 0, 100, 200))

	for p := range 100 * 200 {
		pixelOffset := 4 * p
		myImage.Pix[0+pixelOffset] = uint8(rand.Intn(256)) // Red
		myImage.Pix[1+pixelOffset] = uint8(rand.Intn(256)) // Green
		myImage.Pix[2+pixelOffset] = uint8(rand.Intn(256)) // Blue
		myImage.Pix[2+pixelOffset] = 255                   // Alpha
	}

	outputFile, err := os.Create("test.jpg")
	if err != nil {
		log.Fatal(err)
	}

	jpeg.Encode(outputFile, myImage, nil)

	err = outputFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
