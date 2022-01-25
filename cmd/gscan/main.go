package main

import (
	gscan "github.com/MitchellWT/gscan/internal"
	cli "github.com/MitchellWT/gscan/internal/cli"
)

func main() {
	gscan.Setup()
	cli.Execute()
}
