package gscan

import (
	"fmt"
	"os"
	"testing"

	gscan "github.com/MitchellWT/gscan/internal"
)

const rootDir = "./testingo"

func setUp() {
	os.Mkdir(rootDir, 0755)
}

func tearDown() {
	os.RemoveAll(rootDir)
}

func TestGetAllFilesEmpty(t *testing.T) {
	defer tearDown()
	setUp()

	allFiles := gscan.GetAllFiles(rootDir)
	if len(allFiles) != 0 {
		t.Errorf("Error: len(allFiles) equals %d, should equal 0", len(allFiles))
	}
}

func TestGetAllFilesDeep(t *testing.T) {
	defer tearDown()
	setUp()

	runningDir := rootDir
	for i := 0; i < 100; i++ {
		runningDir += "/herm"
		os.Mkdir(runningDir, 0755)
		os.Create(runningDir + "/epic.md")
	}

	allFiles := gscan.GetAllFiles(rootDir)
	if len(allFiles) != 100 {
		t.Errorf("Error: len(allFiles) equals %d, should equal 100", len(allFiles))
	}
}

func TestGetAllFilesShallow(t *testing.T) {
	defer tearDown()
	setUp()

	for i := 0; i < 100; i++ {
		os.Create(rootDir + fmt.Sprintf("/epic_%d.md", i))
	}

	allFiles := gscan.GetAllFiles(rootDir)
	if len(allFiles) != 100 {
		t.Errorf("Error: len(allFiles) equals %d, should equal 100", len(allFiles))
	}
}
