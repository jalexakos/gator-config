package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/jalexakos/gator-config/internal/database"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

type State struct {
	Db  *database.Queries
	Cfg *Config
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return err
	}
	configFile, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(configFile, jsonBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Read() (*Config, error) {
	// Read config from file or environment variables
	configFile, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	jsonBytes, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(jsonBytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configFileName), nil
}
