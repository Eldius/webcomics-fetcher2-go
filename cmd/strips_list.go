package cmd

import (
	"fmt"

	"github.com/Eldius/webcomics-fetcher2-go/repository"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listStripsCmd = &cobra.Command{
	Use:   "list <plugin-name>",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(_ *cobra.Command, args []string) {
		result, err := repository.NewRepository().ListComicStrip(args[0])
		if err != nil {
			fmt.Printf("Failed to fetch strips: %s\n", err.Error())
		}

		for _, s := range result {
			fmt.Printf("- %s => %s\n", s.Name, s.RelativeFilename())
		}
	},
}

func init() {
	stripsCmd.AddCommand(listStripsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
