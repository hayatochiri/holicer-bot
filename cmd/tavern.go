package cmd

import (
	"github.com/spf13/cobra"
)

var tavernCmd = &cobra.Command{
	Use:   "tavern",
	Short: "Operation on tavern.",
	Long:  `Operation on tavern.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(tavernCmd)
}
