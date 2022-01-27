package main

import (
	gscan "github.com/MitchellWT/gscan/internal"
	"github.com/MitchellWT/gscan/internal/cli"
)

func main() {
	gscan.Setup()
	cli.Execute()
	//gscan.RawExportToHTML("/", "./", enums.All)
}
