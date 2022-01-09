package gscan

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

// GetAllFiles recursively gets all files in the provided root directory
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
			// Check If file is a symlink, ignore
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

// CollectRaw collects raw data from the DataDir directory and returns said
// data in a map[int64][]ScanFile
func CollectRaw(rootDir string, start int64, end int64) map[int64][]ScanFile {
	collectedMap := make(map[int64][]ScanFile)
	files, err := ioutil.ReadDir(DataDir)
	errorCheck(err)

	for _, file := range files {
		fileNameData := strings.Split(file.Name(), "-")
		fileRootDir := strings.ReplaceAll(fileNameData[0], "_", "/")
		fileUnixTime, err := strconv.ParseInt(fileNameData[1][:len(fileNameData[1])-5], 10, 64)
		errorCheck(err)
		// Exit earily for undesirable data files
		if fileRootDir != rootDir || fileUnixTime < start || fileUnixTime > end {
			continue
		}
		fileData, err := os.ReadFile(DataDir + file.Name())
		errorCheck(err)

		fileScanData := ScanData{}
		json.Unmarshal(fileData, &fileScanData)
		collectedMap[fileScanData.UnixTime] = fileScanData.ScanFiles
	}
	return collectedMap
}

// CollectDiff collects diff data from the DataDir directory and returns said
// data in a map[int64][]ScanFile
func CollectDiff(rootDir string, start int64, end int64) map[int64][]ScanFile {
	return nil
}
