package helper

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	SecretKey       string
	Channel         string
	Message         string
	MessagesToEmit  int
	EmitterHostPort string
}

//LoadConfig to load configs from config.yml file
func LoadConfig(fileName string) (*Config, error) {
	var config Config
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
