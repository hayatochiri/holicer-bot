package main

import (
	"fmt"
	"os"

	"./cmd"
	"./holicerBot"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	holicerBot.Initialize()
}
