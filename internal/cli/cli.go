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
	Long: "Allows filesystem metadata collection and aggrigation," +
		"\nData can be output in aggrigated or raw form.",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

var readCmd = &cobra.Command{
	Use:   "read [directory to read]",
	Short: "Reads filesystem metadata and stores it",
	Long: `Reads filesystem name and size and stores this information
				in json format in the /var/lib/gscan/data directory.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		readCommand(cmd, args)
	},
	DisableFlagsInUseLine: true,
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

func init() {
	readCmd.Flags().StringP("out-dir", "o", "", "outputs the current filesystem read in JSON format, to provided dir")

	rootCmd.AddCommand(readCmd)
}

// Execute calls undelying 'Execute' function on the cobra command
func Execute() error {
	return rootCmd.Execute()
}
