package gscan

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	structs "github.com/MitchellWT/gscan/internal/structs"
)

var linuxIgnore []string

func init() {
	linuxIgnore = append(linuxIgnore, "proc", "run", "dev", "sys", "tmp")
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
func GetAllFiles(rootDir string) []structs.ScanFile {
	allFiles := make([]structs.ScanFile, 0)
	return getFiles(rootDir, allFiles)
}

func getFiles(rootDir string, allFiles []structs.ScanFile) []structs.ScanFile {
	files, err := ioutil.ReadDir(rootDir)
	ErrorCheck(err)

	for _, file := range files {
		if file.IsDir() {
			if ignoreSearch(file.Name(), linuxIgnore) {
				continue
			}
			allFiles = getFiles(rootDir+"/"+file.Name(), allFiles)
		} else {
			// Check If file is a symlink, ignore
			if file.Mode()&os.ModeSymlink == os.ModeSymlink {
				continue
			}
			allFiles = append(allFiles, structs.ScanFile{
				Path: rootDir + "/" + file.Name(),
				Size: file.Size(),
			})
		}
	}
	return allFiles
}

// collectRaw collects raw data from the DataDir directory and returns said
// data in a map[int64][]ScanFile
func collectRaw(rootDir string, start int64, end int64) structs.ScanFileMap {
	scanFileMap := make(structs.ScanFileMap)
	files, err := ioutil.ReadDir(LibDir + "data/")
	ErrorCheck(err)

	for _, file := range files {
		fileNameData := strings.Split(file.Name(), "-")
		fileRootDir := strings.ReplaceAll(fileNameData[0], "_", "/")
		fileUnixTime, err := strconv.ParseInt(fileNameData[1][:len(fileNameData[1])-5], 10, 64)
		ErrorCheck(err)
		// Exit earily for undesirable data files
		if fileRootDir != rootDir || fileUnixTime < start || fileUnixTime > end {
			continue
		}
		fileData, err := os.ReadFile(LibDir + "data/" + file.Name())
		ErrorCheck(err)

		fileScanData := structs.ScanData{}
		json.Unmarshal(fileData, &fileScanData)
		scanFileMap[fileScanData.UnixTime] = fileScanData.ScanFiles
	}
	return scanFileMap
}

// collectTotal collects the total size change of the target directory
// from the provided start and end date
func collectTotal(rootDir string, start int64, end int64) structs.TotalDiff {
	totalDiff := make(structs.TotalDiff)
	collectedMap := collectRaw(rootDir, start, end)

	for unixTime, scanSlice := range collectedMap {
		var totalSize int64 = 0
		for _, scanFile := range scanSlice {
			totalSize += scanFile.Size
		}
		totalDiff[unixTime] = totalSize
	}
	return totalDiff
}
