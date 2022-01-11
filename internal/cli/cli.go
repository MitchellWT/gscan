package cli

import (
	gscan "github.com/MitchellWT/gscan/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gscan [command]",
	Short: "Gscan allows users to collect file system metadata",
	Long: "Allows filesystem metadata collection and aggrigation," +
		"\nData can be output in aggrigated or raw form.",
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
	rootDir := args[0]
	allFiles := make([]gscan.ScanFile, 0)
	allFiles = gscan.GetAllFiles(rootDir, allFiles)
	gscan.SaveToJSON(rootDir, gscan.DataDir, allFiles)
}

func init() {
	rootCmd.AddCommand(readCmd)
}

// Execute calls undelying 'Execute' function on the cobra command
func Execute() error {
	return rootCmd.Execute()
}
