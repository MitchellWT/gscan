package gscan

import (
	"io/ioutil"
	"os"
)

var linuxIgnore []string

func init() {
	linuxIgnore = append(linuxIgnore, "proc", "run", "dev", "sys")
}

// Search ignore arrays for passed in string
func ignoreSearch(dirName string, ignoreArr []string) bool {
	for _, ignoreDir := range ignoreArr {
		if ignoreDir == dirName {
			return true
		}
	}
	return false
}

// Recursively gets all files in the provided root directory
func GetAllFiles(rootDir string, allFiles []ScanFile) []ScanFile {
	files, err := ioutil.ReadDir(rootDir)
	errorCheck(err)

	for _, file := range files {
		if file.IsDir() {
			if ignoreSearch(file.Name(), linuxIgnore) {
				continue
			}
			allFiles = GetAllFiles(rootDir+"/"+file.Name(), allFiles)
		} else {
			// Check If file is a Symlink, ignore
			if file.Mode()&os.ModeSymlink == os.ModeSymlink {
				continue
			}
			allFiles = append(allFiles, ScanFile{
				Path: rootDir + "/" + file.Name(),
				Size: file.Size(),
			})
		}
	}
	return allFiles
}
