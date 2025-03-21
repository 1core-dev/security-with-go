// This example uses zip but standard library
// also supports tar archives
package main

import (
	"archive/zip"
	"log"
	"os"
)

func main() {
	// Create a file to write the archive buffer to
	// Could also use an in memory buffer.
	outFile, err := os.Create("test.tar")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Create a zip writer on top of the file writer
	zipWriter := zip.NewWriter(outFile)

	// Adding files to the archive
	// Some hard-coded data is used for demonstration purposes,
	// but the process can involve iterating through all the files
	// in a directory, passing the name and contents of each file,
	// or taking data directly from the program and writing it into
	// the archive.
	var filesToArchive = []struct {
		Name, Body string
	}{
		{"test.txt", "String contents of file"},
		{"test2.txt", "\x61\x62\x63\n"},
	}

	// Create and write files to the archive, which in turn
	// are getting written to the underlying writer to the
	// .zip file we created at the beginning
	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fileWriter.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Clean up
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}
