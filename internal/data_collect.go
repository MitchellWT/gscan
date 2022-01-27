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
func GetAllFiles(rootDir string, allFiles []structs.ScanFile) []structs.ScanFile {
	files, err := ioutil.ReadDir(rootDir)
	ErrorCheck(err)

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
			allFiles = append(allFiles, structs.ScanFile{
				Path: rootDir + "/" + file.Name(),
				Size: file.Size(),
			})
		}
	}
	return allFiles
}

// CollectRaw collects raw data from the DataDir directory and returns said
// data in a map[int64][]ScanFile
func CollectRaw(rootDir string, start int64, end int64) map[int64][]structs.ScanFile {
	collectedMap := make(map[int64][]structs.ScanFile)
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
		collectedMap[fileScanData.UnixTime] = fileScanData.ScanFiles
	}
	return collectedMap
}

// CollectTotal collects the total size change of the target directory
// from the provided start and end date
func CollectTotal(rootDir string, start int64, end int64) map[int64]int64 {
	totalDiff := make(map[int64]int64)
	collectedMap := CollectRaw(rootDir, start, end)

	for unixTime, scanSlice := range collectedMap {
		var totalSize int64 = 0
		for _, scanFile := range scanSlice {
			totalSize += scanFile.Size
		}
		totalDiff[unixTime] = totalSize
	}
	return totalDiff
}
