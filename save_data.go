package gscan

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// SaveToJSON saves passed in data to a file in /var/lib/gscan/data/
func SaveToJSON(rootDir string, scanFiles []ScanFile) {
	currentTime := time.Now().Unix()
	dataFolder := strings.ReplaceAll(rootDir, "/", "_")
	// Builds file name to save data
	fileName := DataDir + dataFolder + "-" + fmt.Sprint(currentTime) + ".json"
	// Builds struct for json storage
	jsonData := ScanData{
		UnixTime:  currentTime,
		RootDir:   rootDir,
		ScanFiles: scanFiles,
	}
	jsonBytes, err := json.Marshal(jsonData)
	errorCheck(err)

	err = os.MkdirAll(DataDir, 0755)
	errorCheck(err)

	err = os.WriteFile(fileName, jsonBytes, 0766)
	errorCheck(err)
}
