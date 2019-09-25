package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

const configDefault = "./api/config/dev.json"

var cfg Config

// Config holds all config info.
type Config struct {
	Service   *Service
}

// Service holds service info.
type Service struct {
	Name        string `json:"name"`
	Port        string `json:"port"`
	ReleaseMode string `json:"release_mode"`
	ApiVersion  string `json:"api_version"`
}


// FromFile reads config from file.
func fromFile(file string) Config {
	var (
		config  Config
		content []byte
		err     error
	)

	content, err = ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func init() {
	var cfgPath string

	flag.StringVar(&cfgPath, "config", configDefault, "Path to config file, by default dev.json")
	flag.Parse()

	cfg = fromFile(cfgPath)
	log.Println("✓ Read configurations from " + cfgPath)
}

// Peek provides secure access to config options.
func Peek() *Config {
	return &cfg
}
