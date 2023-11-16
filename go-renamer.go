package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// path := getPath()
	path := "C:\\Users\\fba95\\Desktop\\root"

	folders := getFiles(path)

	for i := 0; i < len(folders); i++ {
		if folders[i].IsDir() {
			folderPath := filepath.Join(path, folders[i].Name())
			processFolder(folderPath, "")
		}
	}
}

func processFolder(path string, parent string) {
	folderName := filepath.Base(path)

	files := getFiles(path)

	subFolder := getFirstFolder(files)

	for i := 0; i < len(files); i++ {
		if files[i].IsDir() {
			subFolderPath := filepath.Join(path, subFolder.Name())
			processFolder(subFolderPath, folderName)
		} else {
			renameFile(path, parent, folderName, files[i])
		}
	}
	log.Printf("Processed: %s (%d)\n", folderName, len(files))
}

func renameFile(path string, parent string, folderName string, file fs.DirEntry) {
	newName := createNewName(parent, folderName, file.Name())
	os.Rename(filepath.Join(path, file.Name()), filepath.Join(path, newName))
}

func createNewName(parent string, folderName string, fileName string) string {
	if parent == "" {
		return folderName + "-" + fileName
	}
	return parent + "-" + folderName + "-" + fileName
}

func getFirstFolder(files []fs.DirEntry) fs.DirEntry {
	for i := 0; i < len(files); i++ {
		if files[i].IsDir() {
			return files[i]
		}
	}

	return nil
}

func getPath() string {
	if len(os.Args) >= 2 {
		return os.Args[1]
	}
	return getCurrentDirectory()
}

func getFiles(path string) []fs.DirEntry {
	files, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
		os.Exit(3)
	}

	return files
}

func getCurrentDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(3)
	}
	return path
}
