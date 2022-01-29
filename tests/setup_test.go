package test

import (
	"log"
	"os"
	"testing"

	gscan "github.com/MitchellWT/gscan/internal"
)

func TestInitialSetup(t *testing.T) {
	err := os.RemoveAll(gscan.LibDir)
	log.Print(err)

	filesCreated := gscan.Setup()
	if !filesCreated {
		t.Error("Error: gscan.Setup() did not create any files, it should have")
	}
}

func TestPostInitialSetup(t *testing.T) {
	gscan.Setup()

	filesCreated := gscan.Setup()
	if filesCreated {
		t.Error("Error: gscan.Setup() create some files, it shouldn't have")
	}
}
