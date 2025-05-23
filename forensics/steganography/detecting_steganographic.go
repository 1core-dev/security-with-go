package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func main() {
	// Zip signature is “\x50\x4b\x03\x04”
	fileName := "stego_image.jpg"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	bufferedReader := bufio.NewReader(file)

	fileStats, _ := file.Stat()
	// 0 is being cast to an int64 to force i to be initialized as
	// int64 because fileStat.Size() returns an int64 and must be
	// compared against the same type
	for i := range fileStats.Size() {
		myByte, err := bufferedReader.ReadByte()
		if err != nil {
			log.Fatal(err)
		}

		if myByte == '\x50' {
			// First byte match. Check the next 3 bytes
			byteSlice := make([]byte, 3)
			// Get bytes without advancing pointer with Peek
			byteSlice, err = bufferedReader.Peek(3)
			if err != nil {
				log.Fatal(err)
			}

			if bytes.Equal(byteSlice, []byte{'\x4b', '\x03', '\x04'}) {
				log.Printf("Found zip signature at byte %d", i)
			}
		}
	}
}
