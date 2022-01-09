package gscan

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// SaveToFile saves passed in data to a file in /var/lib/gscan/data/
func SaveToFile(rootDir string, scanFiles []ScanFile) {
	currentTime := fmt.Sprint(time.Now().Unix())
	dataFolder := strings.ReplaceAll(rootDir, "/", "_")
	dataDir := "/var/lib/gscan/data/"
	// Builds file name to save data
	fileName := dataDir + dataFolder + "-" + currentTime + ".json"
	// Build struct for json storage
	jsonData := ScanData{
		DateTime:  currentTime,
		RootDir:   rootDir,
		ScanFiles: scanFiles,
	}
	jsonBytes, err := json.Marshal(jsonData)
	errorCheck(err)

	err = os.MkdirAll(dataDir, 0755)
	errorCheck(err)

	err = os.WriteFile(fileName, jsonBytes, 0766)
	errorCheck(err)
}
