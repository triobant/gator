package config

import (
    "encoding/json"
    "os"
    "path/filepath"
)

type Config struct {
    DBURL           string  `json:"db_url"`
    CurrentUserName string  `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {
    cfg.CurrentUserName = userName
    return write(*cfg)
}

func Read() (Config, error) {
    fullPath, err := getConfigFilePath()
    if err != nil {
        return Config{}, err
    }

    file, err := os.Open(fullPath)
    if err != nil {
        return Config{}, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    cfg := Config{}
    err = decoder.Decode(&cfg)
    if err != nil {
        return Config{}, err
    }

    return cfg, nil
}
