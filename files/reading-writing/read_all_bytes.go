package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// Open file for reading
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// os.File.Read(), io.ReadFull(), and
	// io.ReadAtLeast() all work with a fixed
	// byte slice that you make before you read

	// io.ReadAll() will read every byte
	// from the reader (in this case a file),
	// and return a slice of unknown slice
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Data as hex: %x\n", data)
	log.Printf("Data as string: %s", data)
	log.Println("Number of bytes read:", len(data))
}
