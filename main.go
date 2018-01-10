package main

import (
	"fmt"
	"os"

	"github.com/hayatochiri/holicer-bot/cmd"
	"github.com/hayatochiri/holicer-bot/holicerBot"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	holicerBot.Initialize()

	holicerBot.Finalize()
}
