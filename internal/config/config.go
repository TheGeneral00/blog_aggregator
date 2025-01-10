package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Config struct{
    DbURL string `json:"db_url"`
    CurrentUserName string `json:"current_user_name"`
}

/* funtion to read config file. Looks for file at location specified in the os.Open call */
func Read() (*Config) {
        configPath, err := getConfigFilePath()
        if err != nil {
                return nil
        }
        jsonFile, err := os.Open(configPath)
        if err != nil {
                return nil
        }
        defer jsonFile.Close()
        byteContents, err := io.ReadAll(jsonFile)
        if err != nil {
                return nil
        }
        var config *Config
        err = json.Unmarshal(byteContents, &config)
        if err != nil {
                return nil
        }
        if len(config.DbURL) == 0 {
                config.setDBURL()
        }
        return config
}

func (c Config) SetUser(username string) error {
        c.CurrentUserName = username
        configPath, err := getConfigFilePath()
        if err != nil {
                return fmt.Errorf("Failed to retrieve config path with error: %v", err)
    }
        // opens file with write properties and clears the file before writing 
        // sets owner permisson to read and Write
        // only read permisson for other groups 
        file, err := os.OpenFile(configPath, os.O_WRONLY|os.O_TRUNC, 0644)
        if err != nil {
                return fmt.Errorf("Failed to open config filer with error: %v", err)
        }
        defer file.Close()

        data, err := json.Marshal(c)
        if err != nil {
                return fmt.Errorf("Failed to encode config to json with error: %v", err)
        }

        _, err = file.Write(data)
        if err != nil {
                return fmt.Errorf("Failed to write data to config file with error: %v", err)
        }
        return nil
}

func getConfigFilePath() (string, error){
        homeDir, err := os.UserHomeDir()
        if err != nil {
                return "", fmt.Errorf("Failed to fetch hom directory path with error: %v", err)
        }
        configPath := filepath.Join(homeDir, configFileName)
        return configPath, nil
}

func (c Config) setDBURL() {
        c.DbURL = DbURL
}

func (c Config) GetDBURL() string {
        return c.DbURL
}
