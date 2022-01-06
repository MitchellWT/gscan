package gscan

import (
	"io/ioutil"
	"log"
	"os"
)

var linuxIgnore []string

func init() {
	linuxIgnore = append(linuxIgnore, "proc", "run", "dev", "sys")
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ignoreSearch(dirName string, ignoreArr []string) bool {
	for _, ignoreDir := range ignoreArr {
		if ignoreDir == dirName {
			return true
		}
	}
	return false
}

func GetAllFiles(rootFile string, allFiles []ScanFile) []ScanFile {
	files, err := ioutil.ReadDir(rootFile)
	errorCheck(err)

	for _, file := range files {
		if file.IsDir() {
			if ignoreSearch(file.Name(), linuxIgnore) {
				continue
			}
			allFiles = GetAllFiles(rootFile+"/"+file.Name(), allFiles)
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
