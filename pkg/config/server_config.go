package config

import (
	log "github.com/sirupsen/logrus"
)

type ServerConfig struct {
	Name        string            `yaml:"name"`
	Enabled     bool              `yaml:"enabled"`
	Labels      Labels            `yaml:"labels"`
	Logger      LoggerConfig      `yaml:"logger"`
	Command     string            `yaml:"command"`
	Arguments   map[string]string `yaml:"arguments"`
	Environment map[string]string `yaml:"environment"`
}

type Servers []ServerConfig

func (config ServerConfig) Dump() {
	log.WithFields(log.Fields{"name": config.Name, "labels": config.Labels, "command": config.Command}).Info("Config server")

	for key, value := range config.Arguments {
		log.WithFields(log.Fields{"name": config.Name, "argument": key, "value": value}).Info("Config server argument")
	}

	for key, value := range config.Environment {
		log.WithFields(log.Fields{"name": config.Name, "environment": key, "value": value}).Info("Config server environment")
	}
}
