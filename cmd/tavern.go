package cmd

import (
	"github.com/spf13/cobra"
)

type tavernAddOpts struct {
	ja string
	en string
}

var tavern_add_opts tavernAddOpts

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
	tavernAddCmd.Flags().StringVarP(&tavern_add_opts.ja, "name-ja", "j", "", "Name of tavern(Japanese)")
	tavernAddCmd.Flags().StringVarP(&tavern_add_opts.en, "name-en", "e", "", "Name of tavern(English)")

	RootCmd.AddCommand(tavernCmd)
	tavernCmd.AddCommand(tavernAddCmd)
}
