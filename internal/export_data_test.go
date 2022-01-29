package gscan

import (
	"encoding/json"
	"os"
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

func TestRawExportToJSONBasic(t *testing.T) {
	exportDataSetUp()

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

func TestTotalExportToJSONBasic(t *testing.T) {
	exportDataSetUp()
}
