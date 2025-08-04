package main

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const defaultConfigFile = "config.yaml"

var config = struct {
	Port              uint16
	Endpoint          string
	FormEndpoint      string `yaml:"formEndpoint"`
	SignalCliEndpoint string `yaml:"signalCliEndpoint"`
	Account           string
	AllowedRecipients []string `yaml:"allowedRecipients"`
}{}

func readConfig() {
	var configFile string
	flag.StringVar(&configFile, "c", defaultConfigFile, "config file")
	flag.StringVar(&configFile, "config", defaultConfigFile, "config file")
	flag.Parse()

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("Could not open config file: %v", err)
	}
	defer file.Close()

	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalf("Could not parse config file: %v", err)
	}
}
