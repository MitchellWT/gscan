package gscan

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func rawExportToJSON(rootDir string, outDir string, interval Interval) string {
	currentTime := time.Now().Unix()
	intervalStart := interval.GetStart()
	intervalEnd := interval.GetEnd()
	collectedMap := CollectRaw(rootDir, intervalStart, intervalEnd)
	// Builds file name to save data
	fileName := outDir + "export-" + fmt.Sprint(currentTime) + ".json"
	jsonData := ExportRaw{
		UnixStartTime: intervalStart,
		UnixEndTime:   intervalEnd,
		RootDir:       rootDir,
		ScanFiles:     collectedMap,
	}
	jsonBytes, err := json.Marshal(jsonData)
	errorCheck(err)

	err = os.MkdirAll(outDir, 0755)
	errorCheck(err)

	err = os.WriteFile(fileName, jsonBytes, 0766)
	errorCheck(err)

	return fileName
}

func totalRawExportToJSON(rootDir string, outDir string, interval Interval) string {
	currentTime := time.Now().Unix()
	intervalStart := interval.GetStart()
	intervalEnd := interval.GetEnd()
	totalDiff := CollectTotalRaw(rootDir, intervalStart, intervalEnd)
	// Builds file name to save data
	fileName := outDir + "export-" + fmt.Sprint(currentTime) + ".json"
	jsonData := ExportCollectedRaw{
		UnixStartTime: intervalStart,
		UnixEndTime:   intervalEnd,
		RootDir:       rootDir,
		TotalDiff:     totalDiff,
	}
	jsonBytes, err := json.Marshal(jsonData)
	errorCheck(err)

	err = os.MkdirAll(outDir, 0755)
	errorCheck(err)

	err = os.WriteFile(fileName, jsonBytes, 0766)
	errorCheck(err)

	return fileName
}
