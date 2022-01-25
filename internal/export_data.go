package gscan

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
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

	err = os.WriteFile(fileName, jsonBytes, 0644)
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

	err = os.WriteFile(fileName, jsonBytes, 0644)
	ErrorCheck(err)

	return fileName
}

func generateRandomLineColour() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func TotalExportToHTML(rootDir string, outDir string, interval enums.Interval) string {
	currentTime := time.Now().Unix()
	intervalStart := interval.GetStart()
	intervalEnd := interval.GetEnd()
	totalDiff := CollectTotal(rootDir, intervalStart, intervalEnd)
	// Builds file name to save data
	fileName := outDir + "export-" + fmt.Sprint(currentTime) + ".html"
	labelSlice := make([]string, 0)
	dataSlice := make([]float32, 0)

	for unixTime, totalSize := range totalDiff {
		outputTime := time.Unix(unixTime, 0).Format("15:04 02/01/06")
		outputSize := float32(totalSize) / 1024
		labelSlice = append(labelSlice, string(outputTime))
		dataSlice = append(dataSlice, outputSize)
	}

	templateData := TotalHTMLTemplateData{
		Title:       rootDir + " Export Graph",
		GraphLabels: labelSlice,
		DataSets: []HTMLTemplateDataSet{
			HTMLTemplateDataSet{
				Label:      rootDir,
				LineColour: generateRandomLineColour(),
				Data:       dataSlice,
			},
		},
	}

	template, err := template.ParseFiles(LibDir + "templates/template.html")
	ErrorCheck(err)

	exportFile, err := os.Create(fileName)
	ErrorCheck(err)

	err = template.Execute(exportFile, templateData)
	ErrorCheck(err)

	return fileName
}
