package cmd

import (
	"fmt"
	"os"

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
		if tavern_add_opts.ja == "" && tavern_add_opts.en == "" {
			fmt.Fprintf(os.Stderr, "Japanese or English of tavern name is required.")
			cmd.Help()
			os.Exit(1)
		}
	},
}

func init() {
	tavernAddCmd.Flags().StringVarP(&tavern_add_opts.ja, "name-ja", "j", "", "Name of tavern(Japanese)")
	tavernAddCmd.Flags().StringVarP(&tavern_add_opts.en, "name-en", "e", "", "Name of tavern(English)")

	RootCmd.AddCommand(tavernCmd)
	tavernCmd.AddCommand(tavernAddCmd)
}
