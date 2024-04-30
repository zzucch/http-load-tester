package config

import (
	"encoding/json"
	"os"

	"github.com/charmbracelet/log"
)

type Configuration struct {
	RequestURL        string `json:"requestURL"`
	BombardiersAmount int    `json:"bombardiersAmount"`
}

var Config Configuration

func init() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("failed to read config file", "err", err)
	}

	err = json.Unmarshal(file, &Config)
	if err != nil {
		log.Fatal("failed to unmarshal config file", "err", err)
	}
}
