package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	App struct {
		ApiSecret string `yaml:"API_SECRET", envconfig:"API_SECRET"`
		Env       string `yaml:"env", envconfig:"ENV"`
	} `yaml:"app"`
	Server struct {
		Port string `yaml:"port", envconfig:"SERVER_PORT"`
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"DB_USERNAME", envconfig:"DB_USERNAME"`
		Password string `yaml:"DB_PASSWORD", envconfig:"DB_PASSWORD"`
		Host     string `yaml:"DB_HOST", envconfig:"DB_HOST"`
		Port     string `yaml:"DB_PORT", envconfig:"DB_PORT"`
		Database string `yaml:"DB_DATABASE", envconfig:"DB_DATABASE"`
	} `yaml:"database"`
	Redis struct {
		Password string `yaml:"REDIS_PASSWORD", envconfig:"REDIS_PASSWORD"`
		Host     string `yaml:"REDIS_HOST", envconfig:"REDIS_HOST"`
		Port     string `yaml:"REDIS_PORT", envconfig:"REDIS_PORT"`
		DB       int    `yaml:"REDIS_DATABASE", envconfig:"REDIS_DATABASE"`
	} `yaml:"redis"`
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
