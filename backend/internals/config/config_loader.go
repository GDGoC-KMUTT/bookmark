package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
	"strings"
)

var Env *Config

func init() {
	BootConfiguration()
}

func BootConfiguration() {
	// Attempt to determine the current working directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		logrus.Fatal("[CONFIG] Could not determine the caller's directory")
	}

	dir := filepath.Dir(filename)

	// Add config paths for multiple locations
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	// for main.go and in CI/CD
	viper.AddConfigPath(".")

	// for unit test
	viper.AddConfigPath(filepath.Join(dir, "..", ".."))

	// Attempt to read the config file
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("[CONFIG] Error reading config file: %v", err)
	}

	// Set environment variables and replace dots with underscores
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Unmarshal config into the global Env variable
	if err := viper.Unmarshal(&Env); err != nil {
		logrus.Fatalf("[CONFIG] Error unmarshaling config: %v", err)
	}
	logrus.Infof("[CONFIG] Loaded configuration successfully.")
}
