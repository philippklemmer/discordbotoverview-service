package main

import (
	"fmt"
	"os"

	"github.com/philippklemmer/discordbotoverview-service/bot"
	"github.com/philippklemmer/discordbotoverview-service/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	bot.Start()

	<-make(chan struct{})
	return
}
