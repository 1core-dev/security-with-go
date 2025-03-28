package main

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type FileNode struct {
	FullPath string
	Info     os.FileInfo
}

func insertSorted(fileList *list.List, fileNode FileNode) {
	if fileList.Len() == 0 {
		// If list is empty, just insert and return
		fileList.PushFront(fileNode)
		return
	}

	for element := fileList.Front(); element != nil; element = element.Next() {
		if fileNode.Info.Size() < element.Value.(FileNode).Info.Size() {
			fileList.InsertBefore(fileNode, element)
			return
		}
	}
	fileList.PushBack(fileNode)
}

func GetFilesInDirRecursivelyBySize(fileList *list.List, path string) {
	dirFiles, err := os.ReadDir(path)
	if err != nil {
		log.Println("Error reading directory: " + err.Error())
		return
	}
	for _, dirFile := range dirFiles {
		fullpath := filepath.Join(path, dirFile.Name())
		if dirFile.IsDir() {
			// Recurse into subdirectory
			GetFilesInDirRecursivelyBySize(fileList, fullpath)
		} else if dirFile.Type().IsRegular() {
			// Use os.Stat to get os.FileInfo for regular files
			fileInfo, err := os.Stat(fullpath)
			if err != nil {
				log.Println("Error getting file info for:", fullpath, err)
				continue
			}
			// Now we can pass the os.FileInfo to the FileNode
			insertSorted(fileList, FileNode{FullPath: fullpath, Info: fileInfo})
		}
	}
}

func main() {
	fileList := list.New()
	GetFilesInDirRecursivelyBySize(fileList, "/Users/space/sources/playground/security-with-go")
	for element := fileList.Front(); element != nil; element = element.Next() {
		fmt.Printf("%d ", element.Value.(FileNode).Info.Size())
		fmt.Printf("%s\n", element.Value.(FileNode).FullPath)
	}
}
