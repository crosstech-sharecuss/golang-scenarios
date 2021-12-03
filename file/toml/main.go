package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

func main() {
	appConfig := ReadAppConfigFile()

	fmt.Println(appConfig.Database)
}

func ReadAppConfigFile() *AppConfig {
	tomlFile, err := ioutil.ReadFile("config.toml")
	if err != nil {
		log.Printf("tomlFile.Get err   #%v ", err)
	}

	var config AppConfig

	err = toml.Unmarshal(tomlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &config
}

// Structs that represent config file structure

type AppConfig struct {
	Database DatabaseConfig `toml:"database"`
}

type DatabaseConfig struct {
	Ip       string `toml:"ip"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}
