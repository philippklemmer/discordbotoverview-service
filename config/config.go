package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	// Token is the ID for the Discord Bot
	Token string
	// BotPrefix has to be defined, otherwise the bot will not work
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

// ReadConfig  read the config defined in config.json
func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("config/config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
