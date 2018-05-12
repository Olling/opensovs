package main

import (
	"flag"
	"github.com/olling/slog"
)

type DatabaseConfiguration struct {
	Host string
	Port int
	User string
	Password string
	DatabaseName string
}

type Configuration struct {
	DatabaseConf DatabaseConfiguration
	ApiPort int
}

var (
	ConfigurationPath = flag.String("configurationpath","/etc/opensovs/configuration.json","(Optional) The path to the configuration file")
	Conf Configuration
)

func InitializeConfiguration() {
	slog.PrintDebug("Initializing configuration")
	flag.Parse()
	slog.PrintDebug("Configuration Path: " + *ConfigurationPath)
	ReadJsonFile(*ConfigurationPath, &Conf)
}
