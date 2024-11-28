package config

import (
	"encoding/json"
	"os"
)

type Configuration interface {
}

type config struct {
}

func LoadConfigurations(path string) (Configuration, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var confg *config
	json.Unmarshal(b, confg)
	return confg, nil
}
