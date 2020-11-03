package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	App struct {
		Env string `yaml:"env", envconfig:"ENV"`
	} `yaml:"app"`
	Server struct {
		Port string `yaml:"port", envconfig:"SERVER_PORT"`
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user", envconfig:"DB_USERNAME"`
		Password string `yaml:"pass", envconfig:"DB_PASSWORD"`
		Host     string `yaml:"host", envconfig:"DB_HOST"`
		Port     string `yaml:"port", envconfig:"DB_PORT"`
		Database string `yaml:"provider", envconfig:"DB_DATABASE"`
	} `yaml:"provider"`
}

func Load(env string) (cfg Config, err error) {
	f, err := os.Open(fmt.Sprint(env + ".yml"))
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	return
}
