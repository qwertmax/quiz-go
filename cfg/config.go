// Copyright © 2014, 2015 Maxim Tishchenko.
// All Rights Reserved.

package cfg

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	APP_PORT    string `yaml:"app_port"`
	DB_USERNAME string `yaml:"db_username"`
	DB_PASSWORD string `yaml:"db_password"`
	DB_NAME     string `yaml:"db_name"`
	DB_ADDRESS  string `yaml:"db_address"`
	DB_SSLMODE  string `yaml:"db_ssh_mode"`
	DB_PORT     string `yaml:"db_port"`
}

func getConfig(path string) (Config, error) {
	// yamlPath := c.GlobalString("config")
	yamlPath := path
	config := Config{}

	if _, err := os.Stat(yamlPath); err != nil {
		return config, errors.New("config path not valid")
	}

	ymlData, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(ymlData, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// initialize confi, load and perse it, or user defaukr values.
func Init(filename string) Config {
	config, err := getConfig(filename)
	if err != nil {
		log.Fatal(err.Error())
	}

	return config
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
