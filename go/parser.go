package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

// options used while creating the content
type ParseOptions struct {
	directoryPath string
	pointStyle    string
	spaceCount    int
	ignoreList    []string
}

// parse options entered by user
func commandParser(osArgs []string) ParseOptions {
	selectedOptions := ParseOptions{
		directoryPath: "",
		pointStyle:    "-",
		spaceCount:    3,
		ignoreList: []string{
			".git",
			".DS_Store",
		},
	}
	if len(osArgs) > 0 {
		selectedOptions.directoryPath = osArgs[0]
	} else {
		currentWorkingDirectory, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		selectedOptions.directoryPath = currentWorkingDirectory
		log.Println("Current Working Direcoty: ", currentWorkingDirectory)
	}
	if len(osArgs) > 1 {
		selectedOptions.pointStyle = osArgs[1]
	}
	if len(osArgs) > 2 {
		spaceCount, err := strconv.Atoi(osArgs[2])
		if err != nil {
			// change this with a CONSTANT VARIABLE
			log.Fatal("error while parsing space size")
		} else {
			selectedOptions.spaceCount = spaceCount
			if selectedOptions.spaceCount < 1 {
				selectedOptions.spaceCount = 1
			}
		}
	}
	if len(osArgs) > 3 {
		selectedOptions.ignoreList = append(selectedOptions.ignoreList, osArgs[3:]...)
	}
	return selectedOptions
}

// type sliceOfStrings []string

func Contains(sliceOfStrings []string, searchString string) bool {
	for _, eachString := range sliceOfStrings {
		if eachString == searchString {
			return true
		}
	}
	return false
}

func getContentOfDirectory(basePath string, relativePath string, selectedOption ParseOptions) []string {
	files, err := ioutil.ReadDir(path.Join(basePath, relativePath))
	if err != nil {
		log.Fatal(err)
	}
	output := []string{}
	childDir := []string{}
	for _, f := range files {
		// ignore directory/file mentioned
		// change this to pointer
		if Contains(selectedOption.ignoreList, f.Name()) {
			log.Printf("ignoring %s", f.Name())
		} else {
			currentLine := selectedOption.pointStyle + strings.Repeat(" ", selectedOption.spaceCount-1) + "[" + f.Name() + "]"
			if strings.Contains(relativePath, " ") || strings.Contains(f.Name(), " ") {
				currentLine += "(<" + path.Join(relativePath, f.Name()) + ">)"
			} else {
				currentLine += "(" + path.Join(relativePath, f.Name()) + ")"
			}
			output = append(output, currentLine)
			if f.IsDir() {
				childDir = getContentOfDirectory(basePath, path.Join(relativePath, f.Name()), selectedOption)
				for _, aPoint := range childDir {
					// change this spacing
					output = append(output, strings.Repeat(" ", selectedOption.spaceCount)+aPoint)
				}
			}
		}
	}
	return output
}

func main() {
	// directoryPath is directory in which contents are tabularized
	fmt.Println(os.Args)
	osArgs := os.Args[1:]
	selectedOptions := commandParser(osArgs)
	log.Println(selectedOptions, osArgs)
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Current Working Direcoty: ", currentWorkingDirectory)
	for _, files := range getContentOfDirectory(selectedOptions.directoryPath, "/", selectedOptions) {
		fmt.Println(files)
	}
}
