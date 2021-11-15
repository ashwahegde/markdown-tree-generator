package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mdtreegen/githubutils"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
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

func extractName(inpString string) string {
	_, fileOrDirName := path.Split(inpString)
	return fileOrDirName
}

func getContentOfDirectory(basePath string, relativePath string, selectedOption ParseOptions) []string {
	files, err := ioutil.ReadDir(path.Join(basePath, relativePath))
	if err != nil {
		log.Fatal(err)
	}
	output := []string{}
	// childDir := []string{}
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
				childDir := getContentOfDirectory(basePath, path.Join(relativePath, f.Name()), selectedOption)
				for _, aPoint := range childDir {
					// change this spacing
					output = append(output, strings.Repeat(" ", selectedOption.spaceCount)+aPoint)
				}
			}
		}
	}
	return output
}

func getContentOfRepo(repoTreeUrl string, selectedOption ParseOptions) []string {
	fmt.Println(repoTreeUrl)
	output := []string{}
	aTree := githubutils.GetChildren(repoTreeUrl)
	for _, aNode := range aTree {
		fileOrDirName := extractName(aNode.Path)
		completePath := path.Join("/", aNode.Path)
		if Contains(selectedOption.ignoreList, fileOrDirName) {
			log.Printf("ignoring %s", aNode.Path)
		} else {
			currentLine := selectedOption.pointStyle + strings.Repeat(" ", selectedOption.spaceCount-1) + "[" + fileOrDirName + "]"
			if strings.Contains(completePath, " ") {
				currentLine += "(<" + completePath + ">)"
			} else {
				currentLine += "(" + completePath + ")"
			}
			output = append(output, currentLine)
			if aNode.RType == "tree" {
				childDir := getContentOfRepo(aNode.Url, selectedOption)
				for _, aPoint := range childDir {
					// change this spacing
					output = append(output, strings.Repeat(" ", selectedOption.spaceCount)+aPoint)
				}
			}
		}
	}
	return output
}

func getContentOfRepoParallel(repoTreeUrl string, selectedOption ParseOptions, prevResult chan map[string][]string, prevWg *sync.WaitGroup, prevSha string, prefPath string) {
	fmt.Println(repoTreeUrl)
	output := map[string]string{}
	aTree := githubutils.GetChildren(repoTreeUrl)
	result := make(chan map[string][]string)
	var wg sync.WaitGroup
	for _, aNode := range aTree {
		fileOrDirName := extractName(aNode.Path)
		completePath := path.Join(prefPath, aNode.Path)
		if Contains(selectedOption.ignoreList, fileOrDirName) {
			log.Printf("ignoring %s", aNode.Path)
		} else {
			currentLine := selectedOption.pointStyle + strings.Repeat(" ", selectedOption.spaceCount-1) + "[" + fileOrDirName + "]"
			if strings.Contains(completePath, " ") {
				currentLine += "(<" + completePath + ">)"
			} else {
				currentLine += "(" + completePath + ")"
			}
			output[aNode.Sha] = currentLine
			if aNode.RType == "tree" {
				wg.Add(1)
				go getContentOfRepoParallel(aNode.Url, selectedOption, result, &wg, aNode.Sha, completePath)
			}
		}
	}
	outList := []string{}
	// Close channel after goroutines complete.
	go func() {
		wg.Wait()
		close(result)
	}()
	// get all results
	resultObt := map[string][]string{}
	for childDir := range result {
		for sha, allSubLines := range childDir {
			resultObt[sha] = allSubLines
		}
	}
	// iterate over each childs
	for sha, stringSlice := range output {
		outList = append(outList, stringSlice)
		if val, ok := resultObt[sha]; ok {
			for _, aPoint := range val {
				outList = append(
					outList,
					strings.Repeat(" ", selectedOption.spaceCount)+aPoint,
				)
			}
		}
	}
	finalOutput := map[string][]string{}
	finalOutput[prevSha] = outList
	prevResult <- finalOutput
	prevWg.Done()
}

func main() {
	// directoryPath is directory in which contents are tabularized
	osArgs := os.Args[1:]
	selectedOptions := commandParser(osArgs)
	log.Println(selectedOptions, osArgs)

	if strings.Contains(selectedOptions.directoryPath, "https://") {
		// // serial
		// inpUrl, err := url.Parse(selectedOptions.directoryPath)
		// if err != nil {
		// 	log.Fatal("Invalid repository URL")
		// 	os.Exit(1)
		// }
		// fmt.Println(inpUrl.Path)
		// for _, files := range getContentOfRepo("https://api.github.com/repos"+inpUrl.Path+"/git/trees/master", selectedOptions) {
		// 	fmt.Println(files)
		// }
		// Parallel
		inpUrl, err := url.Parse(selectedOptions.directoryPath)
		if err != nil {
			log.Fatal("Invalid repository URL")
			os.Exit(1)
		}
		fmt.Println(inpUrl.Path)
		result := make(chan map[string][]string)
		var wg sync.WaitGroup
		wg.Add(1)
		go getContentOfRepoParallel("https://api.github.com/repos"+inpUrl.Path+"/git/trees/master", selectedOptions, result, &wg, "1234", "/")
		go func() {
			wg.Wait()
			close(result)
		}()
		fmt.Println(" ")
		for files := range result {
			// fmt.Println(files)
			for _, file := range files {
				for _, eachFile := range file {
					fmt.Println(eachFile)
				}
			}
		}

	} else {
		for _, files := range getContentOfDirectory(selectedOptions.directoryPath, "/", selectedOptions) {
			fmt.Println(files)
		}
	}
}
