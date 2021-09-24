package cmd

import (
	"github.com/Eldius/webcomics-fetcher2-go/server"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts app in web mode",
	Long: `Starts app in web mode.
For example:

webcomic start -p 8000

`,
	Run: func(_ *cobra.Command, _ []string) {
		server.Start(startPort)
	},
}

var (
	startPort int
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().IntVarP(&startPort, "port", "p", 8000, "Port to listen for web app mode: -p 8080")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
