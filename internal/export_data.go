package gscan

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	enums "github.com/MitchellWT/gscan/internal/enums"
)

func RawExportToJSON(rootDir string, outDir string, interval enums.Interval) string {
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
	ErrorCheck(err)

	err = os.MkdirAll(outDir, 0755)
	ErrorCheck(err)

	err = os.WriteFile(fileName, jsonBytes, 0766)
	ErrorCheck(err)

	return fileName
}

func TotalExportToJSON(rootDir string, outDir string, interval enums.Interval) string {
	currentTime := time.Now().Unix()
	intervalStart := interval.GetStart()
	intervalEnd := interval.GetEnd()
	totalDiff := CollectTotal(rootDir, intervalStart, intervalEnd)
	// Builds file name to save data
	fileName := outDir + "export-" + fmt.Sprint(currentTime) + ".json"
	jsonData := ExportCollectedRaw{
		UnixStartTime: intervalStart,
		UnixEndTime:   intervalEnd,
		RootDir:       rootDir,
		TotalDiff:     totalDiff,
	}
	jsonBytes, err := json.Marshal(jsonData)
	ErrorCheck(err)

	err = os.MkdirAll(outDir, 0755)
	ErrorCheck(err)

	err = os.WriteFile(fileName, jsonBytes, 0766)
	ErrorCheck(err)

	return fileName
}

func TotalExportToHTML(rootDir string, outDir string, interval enums.Interval) string {
	// currentTime := time.Now().Unix()
	intervalStart := interval.GetStart()
	intervalEnd := interval.GetEnd()
	totalDiff := CollectTotal(rootDir, intervalStart, intervalEnd)
	// Builds file name to save data
	// fileName := outDir + "export-" + fmt.Sprint(currentTime) + ".html"
	fmt.Println(totalDiff)
	return ""
}
