package main

import (
	"fmt"
	"os"

	"github.com/hayatochiri/holicer-bot/cmd"
	"github.com/hayatochiri/holicer-bot/holicerBot"
)

func main() {
	holicerBot.Initialize()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	holicerBot.Finalize()

	os.Exit(0)
}
