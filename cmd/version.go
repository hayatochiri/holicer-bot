package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of holicer-bot",
	Long:  `All software has versions. This is holicer-bot's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("holicer-bot v0.1.0")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
