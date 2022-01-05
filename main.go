package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// ScanFile is the singular file scanned from the application
type ScanFile struct {
	Path string
	Size int64
}

// ScanData is the data that will be stored in json form in the /var/lib/gscan/data dir
type ScanData struct {
	DateTime  string
	RootDir   string
	ScanFiles []ScanFile
}

var linuxIgnore []string

func ignoreSearch(dirName string, ignoreArr []string) bool {
	for _, ignoreDir := range ignoreArr {
		if ignoreDir == dirName {
			return true
		}
	}
	return false
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getAllFiles(rootFile string, allFiles []ScanFile) []ScanFile {
	files, err := ioutil.ReadDir(rootFile)
	errorCheck(err)

	for _, file := range files {
		if file.IsDir() {
			if ignoreSearch(file.Name(), linuxIgnore) {
				continue
			}
			allFiles = getAllFiles(rootFile+"/"+file.Name(), allFiles)
		} else {
			if file.Mode()&os.ModeSymlink == os.ModeSymlink {
				continue
			}
			allFiles = append(allFiles, ScanFile{
				Path: rootFile + "/" + file.Name(),
				Size: file.Size(),
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

func saveToFile(rootFile string, allFiles []ScanFile) {
	currentTime := time.Now().Format(time.UnixDate)
	dataFolder := strings.ReplaceAll(rootFile, "/", "_")
	dataDir := "/var/lib/gscan/data/"
	fileName := dataDir + dataFolder + " " + currentTime + ".json"
	jsonData := ScanData{
		DateTime:  currentTime,
		RootDir:   rootFile,
		ScanFiles: allFiles,
	}
	jsonBytes, err := json.Marshal(jsonData)
	errorCheck(err)
	err = os.MkdirAll(dataDir, 0755)
	errorCheck(err)
	fmt.Println(string(jsonBytes))
	err = os.WriteFile(fileName, jsonBytes, 0766)
	errorCheck(err)
}

func init() {
	linuxIgnore = append(linuxIgnore, "proc", "run", "dev", "sys")
}

func main() {
	rootFile := getRuntimeArgs()[0]
	allFiles := make([]ScanFile, 0)
	allFiles = getAllFiles(rootFile, allFiles)
	saveToFile(rootFile, allFiles)
}
