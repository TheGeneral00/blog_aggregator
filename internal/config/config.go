package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct{
    DBURL string `json:"db_url"`
    CurrentUserName string `json:"current_user_name"`
}

/* funtion to read config file. Looks for file at location specified in the os.Open call */
func Read() (Config, error) {
        fullPath, err := getConfigFilePath()
        if err != nil {
                return Config{}, nil
        }
        file, err := os.Open(fullPath)
        if err != nil {
                return Config{}, err
        }
        defer file.Close()
        
        decoder := json.NewDecoder(file)
        config := Config{}
        err = decoder.Decode(&config)
        if err != nil {
                return Config{}, err
        }
        return config, nil
}

func write(config Config) error {
        fullPath, err := getConfigFilePath()
        if err != nil {
                return err
        }
        file, err := os.Create(fullPath)
        if err != nil {
                return err
        }
        defer file.Close()

        encoder := json.NewEncoder(file)
        err = encoder.Encode(config)
        if err != nil {
                return err
        }
        return nil
}

func (config *Config) SetUser(username string) error {
        config.CurrentUserName = username
        return write(*config)
}

func getConfigFilePath() (string, error){
        homeDir, err := os.UserHomeDir()
        if err != nil {
                return "", err
        }
        configPath := filepath.Join(homeDir, configFileName)
        return configPath, nil
}
