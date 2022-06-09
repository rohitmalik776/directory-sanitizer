package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	var files []string = []string{}

	filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() {
			// File encountered
			if matcher(info.Name()) {
				fmt.Println(path)
				files = append(files, path)
			}
		}
		return nil
	})

	var userInput string
	fmt.Println("Do you wish to delete all these files? (y/n)")
	fmt.Scanln(&userInput)
	if userInput == "y" {
		for _, path := range files {
			err = os.Remove(path)
			if err != nil {
				fmt.Println("Couldn't delete file:", path)
				fmt.Printf("error: %s", err)
				return
			}
		}
		fmt.Println("Deleted files successfully!")
	} else {
		fmt.Println("Files were not deleted!")
	}
}

func matcher(name string) bool {
	splitName := strings.Split(name, ".")
	if len(splitName) <= 1 {
		return false
	}
	if splitName[1] == "exe" || splitName[1] == "out" {
		return true
	}
	return false
}
