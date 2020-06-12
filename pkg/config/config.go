package config

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

// ServerConfig ..
type ServerConfig struct {
	URL string `yaml:"url"`
}

// CredsConfig ..
type CredsConfig struct {
	APIKEY string `yaml:"apikey"`
}

// Config ...
type Config struct {
	Server ServerConfig `yaml:"server"`
	Creds  CredsConfig  `yaml:"creds"`
}

// GetConfig ...
func GetConfig() (*Config, error) {
	configFile, err := findConfig()
	if err != nil {
		return nil, err
	}

	// Decodes the config file
	var cfg Config
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	if err := validateConfig(cfg); err != nil {
		/* here we would maybe have a helper function that prints out error messages
		   depending on the type of the error */
		return nil, err
	}

	return &cfg, nil
}

func createConfig(directory string) error {
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}

	c := Config{}
	d, err := yaml.Marshal(&c)

	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	// Write bytes to file
	err = ioutil.WriteFile(directory+"/config.yml", []byte(d), 0644)
	if err != nil {
		return err
	}

	log.Print("Created config file please open ~/.config/share-cli/config.yml")
	log.Print("URL: The server API URL (https://example.com/api/v1/upload")
	log.Print("APIKEY: Your API key for using for auth")
	return nil
}

func findConfig() (*os.File, error) {
	// This finds the config if none then creates it.
	directory, err := homedir.Expand("~/.config/share-cli")
	if err != nil {
		return nil, err
	}
	filename := directory + "/config.yml"
	f, err := os.Open(filename)

	if err == nil {
		return f, nil
	}
	if err := createConfig(directory); err != nil {
		return nil, err
	}
	return f, nil
}

func validateConfig(cfg Config) error {
	if cfg.Server.URL == "" && cfg.Creds.APIKEY == "" {
		return errors.New("config file must have the server url and api key")
	}
	if cfg.Server.URL == "" {
		return errors.New("config file must have the server url")
	}
	if cfg.Creds.APIKEY == "" {
		return errors.New("config file must have api key")
	}
	return nil
}
