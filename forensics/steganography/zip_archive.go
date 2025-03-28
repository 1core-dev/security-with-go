package main

import (
	"archive/zip"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"os"
)

func printUsage() {
	fmt.Println("Usage: " + os.Args[0] + " <filepath>")
	fmt.Println("Example: " + os.Args[0] + " document.txt")
}

func checkArgs() string {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	return os.Args[1]
}

func main() {
	filename := checkArgs()

	// Get bytes from file
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Create a zip writer on top of the file writer
	zipWriter := zip.NewWriter(outFile)

	addFileToZip := func(fileName string) error {
		// Open the file to be added to the zip
		file, err := os.Open(fileName)
		if err != nil {
			return fmt.Errorf("error opening file %s: %v", fileName, err)
		}
		defer file.Close()

		// Create a new file header for the ZIP file
		writer, err := zipWriter.Create(fileName)
		if err != nil {
			return fmt.Errorf("error creating zip entry for %s: %v", fileName, err)
		}

		// Copy the content of the file to the zip archive
		_, err = io.Copy(writer, file)
		if err != nil {
			return fmt.Errorf("error writing file to zip: %v", err)
		}

		return nil
	}

	// Add test.txt and test2.txt to the ZIP archive
	files := []string{"test.txt", "test2.txt"}
	for _, fileName := range files {
		err := addFileToZip(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Hash the file and output results
		fmt.Printf("md5: %x\n\n", md5.Sum(data))
		fmt.Printf("sha1: %x\n\n", sha1.Sum(data))
		fmt.Printf("sha256: %x\n\n", sha256.Sum256(data))
		fmt.Printf("sha512: %x\n\n", sha512.Sum512(data))
	}
}
