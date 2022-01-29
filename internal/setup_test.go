package gscan

import (
	"log"
	"os"
	"testing"
)

func TestInitialSetup(t *testing.T) {
	err := os.RemoveAll(LibDir)
	if err != nil {
		log.Print(err)
	}

	filesCreated := Setup()
	if !filesCreated {
		t.Error("Error: Setup() did not create any files, it should have")
	}
}

func TestPostInitialSetup(t *testing.T) {
	Setup()

	filesCreated := Setup()
	if filesCreated {
		t.Error("Error: Setup() create some files, it shouldn't have")
	}
}
