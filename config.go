package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	SkillID      string
	Socket       string
	Port         string
	AriaEndpoint string
	AriaToken    string
}

func NewConfig(configFile string) *Config {
	if configFile == "" {
		configFile = "config.json"
	}

	raw, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Failed to open config file: %s\n", configFile)
	}

	config := new(Config)
	err = json.Unmarshal(raw, config)
	if err != nil {
		log.Fatalln("Failed to parse config json")
	}

	log.Println("SkillID:", config.SkillID)
	log.Println("Socket:", config.Socket)
	log.Println("Port:", config.Port)
	log.Println("AriaEndpoint:", config.AriaEndpoint)
	return config
}

func (c Config) ValidClientID(skillID string) bool {
	if c.SkillID == "" {
		return true
	}

	return skillID == c.SkillID
}
