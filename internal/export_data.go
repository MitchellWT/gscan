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
	currentTime := time.Now().Unix()
	intervalStart := interval.GetStart()
	intervalEnd := interval.GetEnd()
	totalDiff := CollectTotal(rootDir, intervalStart, intervalEnd)
	// Builds file name to save data
	fileName := outDir + "export-" + fmt.Sprint(currentTime) + ".html"
	labelString := "["
	dataString := "["

	for unixTime, totalSize := range totalDiff {
		outputTime := time.Unix(unixTime, 0).Format("15:04 02/01/06")
		outputSize := float32(totalSize) / 1024
		labelString += outputTime + ", "
		sizeString := fmt.Sprintf("%.4f, ", outputSize)
		dataString += sizeString
	}

	labelString = string(labelString[:len(labelString)-2]) + "]"
	dataString = string(dataString[:len(dataString)-2]) + "]"

	// TODO: add HTML template using template.html file, will need
	// to sub in values calculated above, and write to file

	return fileName
}
