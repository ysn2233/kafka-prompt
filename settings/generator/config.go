package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/c-bata/go-prompt"
)

type Config struct {
	Commands        map[string]Command
	CommandsSuggest []prompt.Suggest
	OptionSuggest   map[string][]prompt.Suggest
	BsCommand       map[string]bool
	BlCommand       map[string]bool
	ZkCommand       map[string]bool
}

type Option struct {
	Name        string
	Description string
}

type Command struct {
	Name        string
	Description string
	Options     map[string]Option
}

func NewConfig(filename string) Config {
	var c Config
	err := c.ReadCommandOption(filename)
	if err != nil {
		log.Panic(err)
	}
	c.OptionSuggest = make(map[string][]prompt.Suggest)
	c.BsCommand = make(map[string]bool)
	c.BlCommand = make(map[string]bool)
	c.ZkCommand = make(map[string]bool)
	for command, commandDetail := range c.Commands {
		c.CommandsSuggest = append(c.CommandsSuggest, prompt.Suggest{Text: command, Description: commandDetail.Description})
		for option, optionDetail := range commandDetail.Options {
			c.OptionSuggest[command] = append(c.OptionSuggest[command], prompt.Suggest{Text: option, Description: optionDetail.Description})
			if option == "--bootstrap-server" {
				c.BsCommand[command] = true
			} else if option == "--broker-list" {
				c.BlCommand[command] = true
			} else if option == "--zookeeper" {
				c.ZkCommand[command] = true
			}
		}
	}
	return c
}

func (c *Config) ReadCommandOption(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &c.Commands)
	if err != nil {
		return err
	}
	return nil
}
