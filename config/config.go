package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string
	RolesChannelID string
	RoleIDs []EmojiRoleID 

	config *configStruct
)
type EmojiRoleID struct{
	Emoji string `json:"Emoji"`
	RoleID string `json:"RoleID"`
}
type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
	RolesChannelID string `json:"RolesChannelID"`

	RoleIDs []EmojiRoleID `json:"RoleIDs"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	RolesChannelID = config.RolesChannelID
	RoleIDs = config.RoleIDs
	
	return nil
}