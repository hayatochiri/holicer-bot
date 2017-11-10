package main

import (
	. "./src/key"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.BindEnv("TWITTER_CONSUMER_KEY")
	viper.BindEnv("TWITTER_CONSUMER_SECRET")
	viper.BindEnv("TWITTER_ACCESS_TOKEN")
	viper.BindEnv("TWITTER_ACCESS_TOKEN_SECRET")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Loading config file error: %s \n", err))
	}

	api := GetTwitterApi(
		viper.Get("TWITTER_CONSUMER_KEY").(string),
		viper.Get("TWITTER_CONSUMER_SECRET").(string),
		viper.Get("TWITTER_ACCESS_TOKEN").(string),
		viper.Get("TWITTER_ACCESS_TOKEN_SECRET").(string),
	)

	dm, err := api.GetDirectMessages(nil)
	if err != nil {
		panic(err)
	}

	for i, v := range dm {
		fmt.Print(i)
		fmt.Println(v)
	}
}
