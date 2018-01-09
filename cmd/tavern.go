package cmd

import (
	"github.com/spf13/cobra"
)

var tavernCmd = &cobra.Command{
	Use:   "tavern",
	Short: "Operation on tavern.",
	Long:  `Operation on tavern.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var tavernAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add tavern to database.",
	Long:  `Add tavern to database.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(tavernCmd)
	tavernCmd.AddCommand(tavernAddCmd)
}
