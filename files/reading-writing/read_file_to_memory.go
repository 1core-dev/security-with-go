package main

import (
	"log"
	"os"
)

func main() {
	// os.ReadFile() function will take care of opening, reading, and
	// closing the file.
	data, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Data read: %s\n", data)
}
