package main

import (
	"encoding/json"
	"os"
)

// Config configuration file
type Config struct {
	SharedSecret   string
	Server         string
	Zone           string
	Domain         string
	NsupdateBinary string
	RecordTTL      int
}

// LoadConfig load the config from file
func (conf *Config) LoadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
}
