package reader

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config values
type Config struct {
	Models      string     `yaml:"models"`
	Definitions string     `yaml:"definitions"`
	Connection  Connection `yaml:"connection"`
}

// Connection holds the d connection details
type Connection struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

// impure function
func readConfigFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// pure function
func extractConfigYaml(content []byte, model *Config) (config Config, err error) {
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
