package config

import (
	"encoding/json"
	"os"
)

var (
	Cfg *Configuration = nil
)

type LLMConfig struct {
	Model                string `json:"Model"`
	SystemPromptTemplate string `json:"SystemPromptTemplate"`
	Host                 string `json:"Host"`
	Port                 int    `json:"Port"`
	ChatURL              string `json:"ChatURL"`
	SubmitPattern        string `json:"SubmitPattern"`
}
type Configuration struct {
	LLMConfiguration LLMConfig `json:"LLMConfiguration"`
	Port             int       `json:"Port"`
}

func LoadConfigurations(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var Cfg1 Configuration
	err = json.Unmarshal(b, &Cfg1)
	Cfg = &Cfg1
	return err
}
