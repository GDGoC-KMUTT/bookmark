package config

import (
	"backend/internals/utils"
	"fmt"
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
	utils.BootTimeLocation()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Attempt to read the configuration from the current directory
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Printf("[CONFIG] config.yaml not found in the current path, trying parent directory.")
		// If reading from the current directory fails, try the parent directory
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			logrus.Printf("[CONFIG] Could not get the directory of the current file")
		}

		dir := filepath.Dir(filename)
		configDir := filepath.Join(dir, "..")
		viper.AddConfigPath(configDir)

		// Try reading the configuration from the parent directory
		if err := viper.ReadInConfig(); err != nil {
			logrus.Printf("[CONFIG] config.yaml not found in the current path, trying parent directory.")
		}
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.Unmarshal(&Env); err != nil {
		panic(fmt.Errorf("[CONFIG] fatal loading configuration: %w, maybe due to invalid configuration format", err))
	}

	logrus.Printf("[CONFIG] Loaded Configuration.")
}
