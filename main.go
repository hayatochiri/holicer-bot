package main

import (
	. "./src/key"
	"fmt"
)

func main() {
	api := GetTwitterApi()

	text := "Hello twitter api(GO lang)"
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		panic(err)
	}

	fmt.Print(tweet.Text)
}
