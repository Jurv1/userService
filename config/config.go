package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type PostgresConfig struct {
	Host     string
	Port     int16
	User     string
	Password string
	Database string
}

func (c *PostgresConfig) NewConfig() *PostgresConfig {

	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
