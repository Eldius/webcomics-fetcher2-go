package cmd

import (
	"github.com/Eldius/webcomics-fetcher2-go/plugins"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches strips from webcomic",
	Args:  cobra.ExactArgs(1),
	Long: `Fetches strips from webcomic.
For example:

webcomic fetch <plugin-name>

`,
	Run: func(_ *cobra.Command, args []string) {
		plugins.NewPluginEngine().FetchStrips(args[0])
	},
}

func init() {
	stripsCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
