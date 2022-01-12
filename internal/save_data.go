package gscan

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// SaveToJSON saves passed in data to a file in /var/lib/gscan/data/
func SaveToJSON(rootDir string, outDir string, scanFiles []ScanFile) string {
	currentTime := time.Now().Unix()
	scanFolder := strings.ReplaceAll(rootDir, "/", "_")
	// Builds file name to save data
	fileName := outDir + scanFolder + "-" + fmt.Sprint(currentTime) + ".json"
	// Builds struct for json storage
	jsonData := ScanData{
		UnixTime:  currentTime,
		RootDir:   rootDir,
		ScanFiles: scanFiles,
	}
	jsonBytes, err := json.Marshal(jsonData)
	ErrorCheck(err)

	err = os.MkdirAll(DataDir, 0755)
	ErrorCheck(err)

	err = os.WriteFile(fileName, jsonBytes, 0766)
	ErrorCheck(err)

	return fileName
}
