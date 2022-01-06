package main

import (
	"log"
	"os"

	"github.com/MitchellWT/gscan"
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
	allFiles := make([]gscan.ScanFile, 0)
	allFiles = gscan.GetAllFiles(rootFile, allFiles)
	gscan.SaveToFile(rootFile, allFiles)
}
