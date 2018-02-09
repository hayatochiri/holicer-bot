package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hayatochiri/holicer-bot/holicerBot"
	"github.com/spf13/cobra"
)

type tavernAddOpts struct {
	ja string
	en string
}

type tavernListOpts struct {
	removed bool
}

var tavern_add_opts tavernAddOpts
var tavern_list_opts tavernListOpts

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

		_, err := holicerBot.AddTavern(holicerBot.AddTavernParams{NameJA: tavern_add_opts.ja, NameEN: tavern_add_opts.en})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occurred while addming tavern.")
			os.Exit(1)
		}
	},
}

var tavernListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get list of taverns from database.",
	Long:  `Get list of taverns from database.`,
	Run: func(cmd *cobra.Command, args []string) {
		taverns_list, err := holicerBot.GetTavernsList(holicerBot.GetTavernsListParams{IsRemoved: tavern_list_opts.removed})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occurred while getting list of taverns.")
			os.Exit(1)
		}

		fmt.Println(`ID | Name(ja) | Name(en)`)
		for _, t := range taverns_list {
			fmt.Println(t.Id, `|`, t.NameJA, `|`, t.NameEN)
		}
	},
}

var tavernRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove list of taverns from database.",
	Long:  `Remove list of taverns from database.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			arg_i, _ := strconv.ParseInt(arg, 10, 64)
			is_removed, err := holicerBot.RemoveTavern(arg_i)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error occurred while removeing tavern.")
				os.Exit(1)
			}

			if !is_removed {
				fmt.Fprintf(os.Stderr, "Tavern(ID:%d) has already been removed.", arg_i)
				os.Exit(1)
			}
		}
	},
}

func init() {
	tavernAddCmd.Flags().StringVarP(&tavern_add_opts.ja, "name-ja", "j", "", "Name of tavern(Japanese)")
	tavernAddCmd.Flags().StringVarP(&tavern_add_opts.en, "name-en", "e", "", "Name of tavern(English)")

	tavernListCmd.Flags().BoolVarP(&tavern_list_opts.removed, "removed", "r", false, "Show removed taverns")

	RootCmd.AddCommand(tavernCmd)
	tavernCmd.AddCommand(tavernAddCmd)
	tavernCmd.AddCommand(tavernListCmd)
	tavernCmd.AddCommand(tavernRemoveCmd)
}
