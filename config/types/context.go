package types

import (
	"encoding/json"
	"os"

	"github.com/BurntSushi/toml"
	errors "github.com/pkg/errors"
)

type Config struct {
	Title    string `json:"title"`
	Database struct {
		Host string `json:"host"`
		Port int    `json:"port"`
		DB   string `json:"db"`
		User string `json:"user"`
		Pass string `json:"pass"`
	} `json:"database"`
	Redis struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"redis"`
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
}

func (c *Config) InitSetting() error {
	file, err := os.Open("./config.toml")
	if err != nil {
		return errors.New("config file not found")
	}
	defer file.Close()

	// Decode the config file into a Config struct
	if _, err := toml.NewDecoder(file).Decode(c); err != nil {
		return errors.Wrap(err, "failed to decode config file")
	}

	// Encode the Config struct as JSON for testing purposes
	_, err = json.MarshalIndent(c, "", "  ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal config")
	}
	if c.Database.Port == 0 {
		c.Database.Port = 3306
	}
	if c.Redis.Port == 0 {
		c.Redis.Port = 6379
	}
	return nil
}
