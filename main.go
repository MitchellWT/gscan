package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type scanFile struct {
	path string
	size int64
}

// TODO: get file tranfersal to run in constant time
// TODO: ignore certain linux folders such as /proc
func getAllFiles(rootFile string, allFiles []scanFile) []scanFile {
	files, err := ioutil.ReadDir(rootFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			allFiles = getAllFiles(rootFile+"/"+file.Name(), allFiles)
		} else {
			allFiles = append(allFiles, scanFile{
				rootFile + "/" + file.Name(),
				file.Size(),
			})
		}
	}
	return allFiles
}

func getRuntimeArgs() []string {
	returnArgs := make([]string, 0)
	scanDir := os.Args[1]
	if len(scanDir) > 1 && scanDir[len(scanDir)-1] == '/' {
		scanDir = scanDir[:len(scanDir)-1]
	}
	returnArgs = append(returnArgs, scanDir)
	return returnArgs
}

func main() {
	rootFile := getRuntimeArgs()[0]
	allFiles := make([]scanFile, 0)
	allFiles = getAllFiles(rootFile, allFiles)
	fmt.Println(allFiles)
}
