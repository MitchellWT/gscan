package gscan

import (
	"encoding/json"
	"os"
	"testing"

	structs "github.com/MitchellWT/gscan/internal/structs"
)

func TestSaveToJSONPreserveEmpty(t *testing.T) {
	rootDir := "I/AM/A/TEST/ROOT/DIR"
	fileString := SaveToJSON(rootDir, "./", []structs.ScanFile{})
	data, err := os.ReadFile(fileString)
	ErrorCheck(err)
	os.Remove(fileString)

	jsonData := structs.ScanData{}
	json.Unmarshal(data, &jsonData)

	if jsonData.RootDir != rootDir {
		t.Errorf("Error: rootDir not preserved in json file, rootDir in json file equals %s", jsonData.RootDir)
	}
	if len(jsonData.ScanFiles) != 0 {
		t.Errorf("Error: len(jsonData.ScanFileMap) equals %d, should equal 0", len(jsonData.ScanFiles))
	}
}
