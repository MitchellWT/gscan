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

func readCommand(cmd *cobra.Command, args []string) {
	outDir := cmd.Flag("out-dir").Value.String()
	rootDir := args[0]
	allFiles := make([]gscan.ScanFile, 0)
	allFiles = gscan.GetAllFiles(rootDir, allFiles)
	saveFile := gscan.SaveToJSON(rootDir, gscan.DataDir, allFiles)
	if outDir == "" {
		return
	}
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		os.Remove(saveFile)
		log.Fatal(err)
	}
	gscan.SaveToJSON(rootDir, outDir, allFiles)
}

func exportCommand(cmd *cobra.Command, args []string) {
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
