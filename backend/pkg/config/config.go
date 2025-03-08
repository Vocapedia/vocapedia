package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	GrpcPort      int    `yaml:"grpc_port"`
	Env           string `yaml:"env"`
	BaseURL       string `yaml:"base_url"`
	FrontBaseURL  string `yaml:"front_base_url"`
	JwtSecret     string `yaml:"jwt_secret"`
	JwtExpire     int    `yaml:"jwt_expire"`
	JwtIssuer     string `yaml:"jwt_issuer"`
	TotpIssuer    string `yaml:"totp_issuer"`
	AdminEmail    string `yaml:"admin_email"`
	AdminName     string `yaml:"admin_name"`
	AdminPassword string `yaml:"admin_password"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`

	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`

	AllowMethods []string `yaml:"allow_methods"`
	AllowHeaders []string `yaml:"allow_headers"`
	AllowOrigins []string `yaml:"allow_origins"`
}

var configs *Config

func ReadValue() *Config {
	if configs != nil {
		return configs
	}
	filename, err := filepath.Abs("./config.yaml")
	if err != nil {
		log.Fatal("error loading config.yaml ", err)
	}
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("error loading config.yaml ", err)
	}
	err = yaml.Unmarshal(yamlFile, &configs)
	if err != nil {
		log.Fatal("error loading config.yaml ", err)
	}
	return configs
}
