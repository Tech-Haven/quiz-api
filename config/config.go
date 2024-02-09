package config

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Listen_Addr string `yaml:"listen_addr" env:"LISTEN_ADDR"`
	Mongo       struct {
		Host       string `yaml:"host" env:"MONGO_HOST"`
		Port       string `yaml:"port" env:"MONGO_PORT"`
		User       string `yaml:"user" env:"MONGO_USER"`
		Password   string `yaml:"password" env:"MONGO_PASSWORD"`
		Database   string `yaml:"database" env:"MONGO_DATABASE"`
		Collection string `yaml:"collection" env:"MONGO_COLLECTION"`
	}
}

func New() *Config {
	var cfg Config
	loadConfig(&cfg)
	loadEnv(&cfg)

	return &cfg
}

func loadConfig(cfg *Config) {
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("unable to load config.yaml file: %s", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func loadEnv(cfg *Config) {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("unable to load .env file: %s", err)
		}
		if err := env.Parse(cfg); err != nil {
			log.Fatalf("unable to parse envionment variables: %s", err)
		}

		fmt.Printf("Parsed .env %+v\n", cfg)
	}
}
