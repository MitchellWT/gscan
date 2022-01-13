package cli

import (
	"log"
	"os"

	gscan "github.com/MitchellWT/gscan/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gscan [command]",
	Short: "Gscan allows users to collect file system metadata",
	Long: "Allows filesystem metadata collection and aggrigation, " +
		"\nData can be output in aggrigated or raw form.",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

var readCmd = &cobra.Command{
	Use:   "read [directory to read]",
	Short: "Reads filesystem metadata and stores it",
	Long: "Reads filesystem information (name and size) and stores " +
		"\nthis information in json format in the " +
		"\n/var/lib/gscan/data directory.",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		readCommand(cmd, args)
	},
	DisableFlagsInUseLine: true,
}

var exportCmd = &cobra.Command{
	Use:   "export [read directory to export]",
	Short: "Exports stores filesystem metadata",
	Long: "Exports filesystem information (name and size) in " +
		"\none of the provided file formats (default minified JSON).",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		exportCommand(cmd, args)
	},
}

// checkDir performs some basic checks to ensure that the provided dir path
// is correct, these included checking last rune and If the dir exists
func checkDir(inputDir string) string {
	if rune(inputDir[len(inputDir)-1]) != '/' {
		inputDir = inputDir + "/"
	}
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		log.Fatal(err)
	}
	return inputDir
}

func readCommand(cmd *cobra.Command, args []string) {
	outDir := cmd.Flag("out-dir").Value.String()
	rootDir := checkDir(args[0])
	allFiles := make([]gscan.ScanFile, 0)
	allFiles = gscan.GetAllFiles(rootDir, allFiles)

	if outDir != "" {
		outDir = checkDir(outDir)
		gscan.SaveToJSON(rootDir, outDir, allFiles)
	}
	gscan.SaveToJSON(rootDir, gscan.DataDir, allFiles)
}

func exportCommand(cmd *cobra.Command, args []string) {
	outDir := checkDir(cmd.Flag("out-dir").Value.String())
	interval, err := gscan.ToInterval(cmd.Flag("interval").Value.String())
	gscan.ErrorCheck(err)

	exportType, err := gscan.ToExportType(cmd.Flag("type").Value.String())
	gscan.ErrorCheck(err)

	rootDir := checkDir(args[0])

	switch exportType {
	case gscan.Raw:
		gscan.RawExportToJSON(rootDir, outDir, interval)
	case gscan.Total:
		gscan.TotalRawExportToJSON(rootDir, outDir, interval)
	}
}

func init() {
	readCmd.Flags().StringP("out-dir", "o", "", "outputs the current filesystem read in JSON format, to provided dir")

	exportCmd.Flags().StringP("out-dir", "o", "./", "directory where the exported file should be saved to")
	exportCmd.Flags().StringP("interval", "i", "all", "interval that the export should take data from, this interval "+
		"\ncan be one of the following values: "+
		"\n- hour "+
		"\n- day "+
		"\n- week "+
		"\n- month "+
		"\n- three-months "+
		"\n- six-months "+
		"\n- year "+
		"\n- all")
	exportCmd.Flags().StringP("type", "t", "raw", "export type denotes how the data should be exported, the data "+
		"\ncan be exported in the following ways: "+
		"\n- total: sums all the files in the root directory and returns the total size "+
		"\n- raw: returns the raw data stored in the data directory copiled together")

	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(exportCmd)
}

// Execute calls undelying 'Execute' function on the cobra command
func Execute() error {
	return rootCmd.Execute()
}
