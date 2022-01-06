package gscan

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

func SaveToFile(rootFile string, allFiles []ScanFile) {
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
	err = os.WriteFile(fileName, jsonBytes, 0766)
	errorCheck(err)
}
