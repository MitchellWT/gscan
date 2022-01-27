package gscan

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"time"

	enums "github.com/MitchellWT/gscan/internal/enums"
	structs "github.com/MitchellWT/gscan/internal/structs"
)

func generateRandomLineColour() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func RawExportToJSON(rootDir string, outDir string, interval enums.Interval) string {
	currentTime := time.Now().Unix()
	intervalStart := interval.GetStart()
	intervalEnd := interval.GetEnd()
	scanFileMap := CollectRaw(rootDir, intervalStart, intervalEnd)
	// Builds file name to save data
	fileName := outDir + "export-" + fmt.Sprint(currentTime) + ".json"
	jsonData := structs.ExportRaw{
		UnixStartTime: intervalStart,
		UnixEndTime:   intervalEnd,
		RootDir:       rootDir,
		ScanFileMap:   scanFileMap,
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
	jsonData := structs.ExportCollected{
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

func RawExportToHTML(rootDir string, outDir string, interval enums.Interval) string {
	currentTime := time.Now().Unix()
	intervalStart := interval.GetStart()
	intervalEnd := interval.GetEnd()
	scanFileMap := CollectRaw(rootDir, intervalStart, intervalEnd)
	// Builds file name to save data
	fileName := outDir + "export-" + fmt.Sprint(currentTime) + ".html"
	labelSlice := make([]string, 0)
	dataSetMap := make(map[string]structs.HTMLTemplateDataSet, 0)

	unixTimeforFileMax := int64(0)
	maxFileAmout := 0

	for unixTime, scanFiles := range scanFileMap {
		if len(scanFiles) > maxFileAmout {
			unixTimeforFileMax = unixTime
			maxFileAmout = len(scanFiles)
		}
	}

	for _, scanFile := range scanFileMap[unixTimeforFileMax] {
		initialDataSet := structs.HTMLTemplateDataSet{
			Label:      scanFile.Path,
			LineColour: generateRandomLineColour(),
			Data:       make([]float32, len(scanFileMap)),
		}
		dataSetMap[scanFile.Path] = initialDataSet
	}

	counter := 0

	for unixTime, scanFiles := range scanFileMap {
		outputTime := time.Unix(unixTime, 0).Format("15:04 02/01/06")

		for _, scanFile := range scanFiles {
			dataSetMap[scanFile.Path].Data[counter] = float32(scanFile.Size) / 1024
		}

		labelSlice = append(labelSlice, outputTime)
		counter += 1
	}

	dataSetSlice := make([]structs.HTMLTemplateDataSet, 0)

	for _, DataSet := range dataSetMap {
		dataSetSlice = append(dataSetSlice, DataSet)
	}

	templateData := structs.TotalHTMLTemplateData{
		Title:       rootDir + " Export Graph",
		GraphLabels: labelSlice,
		DataSets:    dataSetSlice,
	}

	template, err := template.ParseFiles(LibDir + "templates/template.html")
	ErrorCheck(err)

	err = os.MkdirAll(outDir, 0755)
	ErrorCheck(err)

	exportFile, err := os.Create(fileName)
	ErrorCheck(err)

	err = template.Execute(exportFile, templateData)
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
	labelSlice := make([]string, 0)
	dataSlice := make([]float32, 0)

	for unixTime, totalSize := range totalDiff {
		outputTime := time.Unix(unixTime, 0).Format("15:04 02/01/06")
		outputSize := float32(totalSize) / 1024
		labelSlice = append(labelSlice, string(outputTime))
		dataSlice = append(dataSlice, outputSize)
	}

	templateData := structs.TotalHTMLTemplateData{
		Title:       rootDir + " Export Graph",
		GraphLabels: labelSlice,
		DataSets: []structs.HTMLTemplateDataSet{
			structs.HTMLTemplateDataSet{
				Label:      rootDir,
				LineColour: generateRandomLineColour(),
				Data:       dataSlice,
			},
		},
	}

	template, err := template.ParseFiles(LibDir + "templates/template.html")
	ErrorCheck(err)

	err = os.MkdirAll(outDir, 0755)
	ErrorCheck(err)

	exportFile, err := os.Create(fileName)
	ErrorCheck(err)

	err = template.Execute(exportFile, templateData)
	ErrorCheck(err)

	return fileName
}

func RawExportToCSV(rootDir string, outDir string, interval enums.Interval) string {
	currentTime := time.Now().Unix()
	intervalStart := interval.GetStart()
	intervalEnd := interval.GetEnd()
	scanFileMap := CollectRaw(rootDir, intervalStart, intervalEnd)
	// Builds file name to save data
	fileName := outDir + "export-" + fmt.Sprint(currentTime) + ".csv"
	csvData := [][]string{
		{"unix_time", "file_path", "file_size"},
	}

	for unixTime, ScanFiles := range scanFileMap {
		for _, ScanFile := range ScanFiles {
			csvData = append(csvData, []string{
				fmt.Sprint(unixTime),
				ScanFile.Path,
				fmt.Sprint(ScanFile.Size),
			})
		}
	}

	err := os.MkdirAll(outDir, 0755)
	ErrorCheck(err)

	exportFile, err := os.Create(fileName)
	ErrorCheck(err)

	exportWriter := csv.NewWriter(exportFile)
	exportWriter.WriteAll(csvData)

	err = exportWriter.Error()
	ErrorCheck(err)

	return fileName
}

func TotalExportToCSV(rootDir string, outDir string, interval enums.Interval) string {
	currentTime := time.Now().Unix()
	intervalStart := interval.GetStart()
	intervalEnd := interval.GetEnd()
	totalDiff := CollectTotal(rootDir, intervalStart, intervalEnd)
	// Builds file name to save data
	fileName := outDir + "export-" + fmt.Sprint(currentTime) + ".csv"
	csvData := [][]string{
		{"unix_time", "total_size"},
	}

	for unixTime, totalSize := range totalDiff {
		csvData = append(csvData, []string{
			fmt.Sprint(unixTime),
			fmt.Sprint(totalSize),
		})
	}

	err := os.MkdirAll(outDir, 0755)
	ErrorCheck(err)

	exportFile, err := os.Create(fileName)
	ErrorCheck(err)

	exportWriter := csv.NewWriter(exportFile)
	exportWriter.WriteAll(csvData)

	err = exportWriter.Error()
	ErrorCheck(err)

	return fileName
}
