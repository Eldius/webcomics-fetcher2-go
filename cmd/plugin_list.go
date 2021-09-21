package cmd

import (
	"fmt"

	"github.com/Eldius/webcomics-fetcher2-go/plugins"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List registered plugins",
	Long:  `List registered plugins.`,
	Run: func(_ *cobra.Command, _ []string) {
		r := plugins.NewPluginEngine().ListRegisteredPlugins()
		fmt.Println("Registered plugins:")
		for _, v := range r {
			fmt.Printf("- %s: %s\n", v.Name, v.Description)
		}
	},
}

func init() {
	pluginCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
