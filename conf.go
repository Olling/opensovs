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
	LogLevel = flag.Int("loglevel",4,"(Optional) The loglevel to apply (from Fatal(1) to Trace(6)")
	Conf Configuration
)

func InitializeConfiguration() {
	slog.PrintDebug("Initializing configuration")
	flag.Parse()
	slog.PrintDebug("Configuration Path: " + *ConfigurationPath)
	slog.SetLogLevel(*LogLevel)
	slog.PrintDebug("Log Level:",*LogLevel)
	ReadJsonFile(*ConfigurationPath, &Conf)
}
