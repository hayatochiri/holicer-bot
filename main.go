package main

import (
	. "./src/key"
	"fmt"
)

func main() {
	api := GetTwitterApi()

	dm, err := api.GetDirectMessages(nil)
	if err != nil {
		panic(err)
	}

	for i, v := range dm {
		fmt.Print(i)
		fmt.Println(v)
	}
}
