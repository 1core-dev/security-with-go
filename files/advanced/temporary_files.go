package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a temp dir in the system default temp folder
	tempDirPath, err := os.MkdirTemp("", "myTempDir")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temp dir created:", tempDirPath)

	// Create a file in new temp directory
	tempFile, err := os.CreateTemp(tempDirPath, "myTempFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temp file created:", tempFile.Name())

	// ... do something with temp file/dir ...

	// Close file
	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Delete the resources we created
	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	}
}
