package cmd

import (
	"github.com/spf13/cobra"
)

var addTavernCmd = &cobra.Command{
	Use:   "add-tavern",
	Short: "Add tavern to the database.",
	Long:  `Add tavern to the database.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(addTavernCmd)
}
