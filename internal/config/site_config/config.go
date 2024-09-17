package site_config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var SiteConfig *Config

// Config represents the configuration structure for the application
type Config struct {
	ServerPort int `json:"server_port"`
	Database   struct {
		URL      string `json:"database_url,omitempty"` // Optional field to maintain compatibility
		Username string `json:"db_user"`
		Password string `json:"db_password"`
		DBName   string `json:"db_name"`
	} `json:"db"` // Embedded struct for database configuration
}

// LoadConfig attempts to load the configuration from the specified filename
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close() // Ensure file is closed even on errors

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
