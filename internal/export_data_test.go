package gscan

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"regexp"
	"testing"

	enums "github.com/MitchellWT/gscan/internal/enums"
	structs "github.com/MitchellWT/gscan/internal/structs"
)

func exportDataSetUp() {
	err := os.MkdirAll(LibDir+"data/", 0755)
	ErrorCheck(err)

	err = os.MkdirAll(LibDir+"templates/", 0755)
	ErrorCheck(err)
}

func TestRawExportToJSONPreserveEmpty(t *testing.T) {
	exportDataSetUp()

	rootDir := "I/AM/A/TEST/ROOT/DIR"
	fileString := RawExportToJSON(rootDir, "./", enums.All)
	data, err := os.ReadFile(fileString)
	ErrorCheck(err)
	os.Remove(fileString)

	jsonData := structs.ExportRaw{}
	json.Unmarshal(data, &jsonData)

	if jsonData.RootDir != rootDir {
		t.Errorf("Error: rootDir not preserved in json file, rootDir in json file equals %s", jsonData.RootDir)
	}
	if len(jsonData.ScanFileMap) != 0 {
		t.Errorf("Error: len(jsonData.ScanFileMap) equals %d, should equal 0", len(jsonData.ScanFileMap))
	}
}

func TestTotalExportToJSONPreserveEmpty(t *testing.T) {
	exportDataSetUp()

	rootDir := "I/AM/A/TEST/ROOT/DIR"
	fileString := TotalExportToJSON(rootDir, "./", enums.All)
	data, err := os.ReadFile(fileString)
	ErrorCheck(err)
	os.Remove(fileString)

	jsonData := structs.ExportTotal{}
	json.Unmarshal(data, &jsonData)

	if jsonData.RootDir != rootDir {
		t.Errorf("Error: rootDir not preserved in json file, rootDir in json file equals %s", jsonData.RootDir)
	}
	if len(jsonData.TotalDiff) != 0 {
		t.Errorf("Error: len(jsonData.TotalDiff) equals %d, should equal 0", len(jsonData.TotalDiff))
	}
}

func TestRawExportToHTMLPreserveEmpty(t *testing.T) {
	exportDataSetUp()

	rootDir := "I/AM/A/TEST/ROOT/DIR"
	fileString := RawExportToHTML(rootDir, "./", enums.All)
	data, err := os.ReadFile(fileString)
	ErrorCheck(err)
	os.Remove(fileString)

	re := regexp.MustCompile(`<h1>.*</h1>`)
	titleSlice := re.FindAllString(string(data), -1)

	re = regexp.MustCompile(`const labels = .*;`)
	labelSlice := re.FindAllString(string(data), -1)

	re = regexp.MustCompile(`data: .[^data]*,`)
	dataSlice := re.FindAllString(string(data), -1)

	if len(titleSlice) != 1 {
		t.Errorf("Error: len(findSlice) equals %d, should equal 1", len(titleSlice))
	}
	if titleSlice[0] != "<h1>I/AM/A/TEST/ROOT/DIR Export Graph</h1>" {
		t.Errorf("Error: findSlice[0] equals %s, should equal '<h1>I/AM/A/TEST/ROOT/DIR Export Graph</h1>'", titleSlice[0])
	}
	if len(labelSlice) != 1 {
		t.Errorf("Error: len(labelSlice) equals %d, should equal 1", len(labelSlice))
	}
	if labelSlice[0] != "const labels = [];" {
		t.Errorf("Error: labelSlice[0] equals %s, should equal 'const labels = [];'", labelSlice[0])
	}
	if len(dataSlice) != 0 {
		t.Errorf("Error: len(dataSlice) equals %d, should equal 0", len(dataSlice))
	}
}

func TestTotalExportToHTMLPreserveEmpty(t *testing.T) {
	exportDataSetUp()

	rootDir := "I/AM/A/TEST/ROOT/DIR"
	fileString := TotalExportToHTML(rootDir, "./", enums.All)
	data, err := os.ReadFile(fileString)
	ErrorCheck(err)
	os.Remove(fileString)

	re := regexp.MustCompile(`<h1>.*</h1>`)
	titleSlice := re.FindAllString(string(data), -1)

	re = regexp.MustCompile(`const labels = .*;`)
	labelSlice := re.FindAllString(string(data), -1)

	re = regexp.MustCompile(`data: .[^data]*,`)
	dataSlice := re.FindAllString(string(data), -1)

	if len(titleSlice) != 1 {
		t.Errorf("Error: len(findSlice) equals %d, should equal 1", len(titleSlice))
	}
	if titleSlice[0] != "<h1>I/AM/A/TEST/ROOT/DIR Export Graph</h1>" {
		t.Errorf("Error: findSlice[0] equals %s, should equal '<h1>I/AM/A/TEST/ROOT/DIR Export Graph</h1>'", titleSlice[0])
	}
	if len(labelSlice) != 1 {
		t.Errorf("Error: len(labelSlice) equals %d, should equal 1", len(labelSlice))
	}
	if labelSlice[0] != "const labels = [];" {
		t.Errorf("Error: labelSlice[0] equals %s, should equal 'const labels = [];'", labelSlice[0])
	}
	if len(dataSlice) != 1 {
		t.Errorf("Error: len(dataSlice) equals %d, should equal 1", len(dataSlice))
	}
	if dataSlice[0] != "data: []," {
		t.Errorf("Error: dataSlice[0] equals %s, should equal 'data: [],'", dataSlice[0])
	}
}

func TestRawExportToCSVEmpty(t *testing.T) {
	exportDataSetUp()

	rootDir := "I/AM/A/TEST/ROOT/DIR"
	fileString := RawExportToCSV(rootDir, "./", enums.All)
	csvFile, err := os.Open(fileString)
	ErrorCheck(err)
	os.Remove(fileString)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	ErrorCheck(err)

	if len(csvLines) != 1 {
		t.Errorf("Error: len(csvLines) equals %d, should equal 1", len(csvLines))
	}
	if csvLines[0][0] != "unix_time" {
		t.Errorf("Error: csvLines[0][0] equals %s, should equal 'unix_time'", csvLines[0][0])
	}
	if csvLines[0][1] != "file_path" {
		t.Errorf("Error: csvLines[0][1] equals %s, should equal 'file_path'", csvLines[0][1])
	}
	if csvLines[0][2] != "file_size" {
		t.Errorf("Error: csvLines[0][2] equals %s, should equal 'file_size'", csvLines[0][2])
	}
}

func TestTotalExportToCSVEmpty(t *testing.T) {
	exportDataSetUp()

	rootDir := "I/AM/A/TEST/ROOT/DIR"
	fileString := TotalExportToCSV(rootDir, "./", enums.All)
	csvFile, err := os.Open(fileString)
	ErrorCheck(err)
	os.Remove(fileString)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	ErrorCheck(err)

	if len(csvLines) != 1 {
		t.Errorf("Error: len(csvLines) equals %d, should equal 1", len(csvLines))
	}
	if csvLines[0][0] != "unix_time" {
		t.Errorf("Error: csvLines[0][0] equals %s, should equal 'unix_time'", csvLines[0][0])
	}
	if csvLines[0][1] != "total_size" {
		t.Errorf("Error: csvLines[0][1] equals %s, should equal 'total_size'", csvLines[0][1])
	}
}
