package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	appConfig := ReadAppConfigFile()

	fmt.Println(appConfig.Database)
}

func ReadAppConfigFile() *AppConfig {
	jsonFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Printf("jsonFile.Get err   #%v ", err)
	}

	var config AppConfig

	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &config
}

// Structs that represent config file structure

type AppConfig struct {
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Ip       string `json:"ip"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}
