package main

import (
	"fmt"
	"os"

	"github.com/MitchellWT/gscan"
)

// Gets the runtime arguments from cli call. Currently gets:
// - root directory to record/save
func getRuntimeArgs() []string {
	scanDir := os.Args[1]
	// Removes trailing '/' from root directory input
	if len(scanDir) > 1 && scanDir[len(scanDir)-1] == '/' {
		scanDir = scanDir[:len(scanDir)-1]
	}

	returnArgs := make([]string, 0)
	returnArgs = append(returnArgs, scanDir)
	return returnArgs
}

func main() {
	// Get root file directory
	//rootFile := getRuntimeArgs()[0]
	//allFiles := make([]gscan.ScanFile, 0)
	// Get all files in root directory
	//allFiles = gscan.GetAllFiles(rootFile, allFiles)
	// Save recorded files to /var/lib/gscan/data/
	//gscan.SaveToFile(rootFile, allFiles)
	collectedMap := gscan.CollectRaw("/home/mitchell/Scripts", 1641721436, 1641741436)
	fmt.Println(collectedMap)
}
