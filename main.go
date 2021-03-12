package main

import (
	"github.com/AUTGamecraft/RoleBot/bot"
	"github.com/AUTGamecraft/RoleBot/config"
	"fmt"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}