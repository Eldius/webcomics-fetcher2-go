package cmd

import (
	"fmt"

	"github.com/Eldius/webcomics-fetcher2-go/plugins"
	"github.com/spf13/cobra"
)

// refreshCmd represents the refresh command
var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh plugin list",
	Long: `Refresh plugin list.
For example:

webcomic refresh

`,
	Run: func(_ *cobra.Command, _ []string) {
		r := plugins.NewPluginEngine().RefreshPluginList()
		fmt.Println("Available plugins:")
		for k, v := range r {
			fmt.Printf("- %s: %s\n", k, v.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// refreshCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// refreshCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
