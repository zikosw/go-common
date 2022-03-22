package envconfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

func read[T any](prefix string, filePath string) (*T, error) {
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error
		} else {
			// Config file was found but another error was produced
			return nil, fmt.Errorf("config file: %w", err)
		}
	}

	for k, v := range viper.AllSettings() {
		if err := os.Setenv(strings.ToUpper(k), fmt.Sprint(v)); err != nil {
			return nil, fmt.Errorf("setenv: %w", err)
		}
	}

	dest := new(T)
	if err := envconfig.Process(prefix, dest); err != nil {
		return nil, fmt.Errorf("process env: %w", err)
	}
	return dest, nil
}

func Read[T any](prefix string) (*T, error) {
	return read[T](prefix, ".env")
}

func ReadWithFile[T any](prefix string, filePath string) (*T, error) {
	return read[T](prefix, filePath)
}
