package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

const (
	// configName is the prefix of environment variables.
	configName = "config"
	envPrefix  = "sstt"
)

var (
	_initConfigOnce sync.Once //nolint:gochecknoglobals
)

// Init initialize config.
func Init() error {
	var err error

	_initConfigOnce.Do(
		func() {
			viper.SetEnvPrefix(envPrefix)
			viper.SetConfigName(configName)
			viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
			viper.SetConfigType("yaml")
			viper.AddConfigPath(".")
			viper.AutomaticEnv()
			err = viper.ReadInConfig()
			viper.WriteConfigAs("config.env")
		},
	)

	if err != nil {
		return fmt.Errorf("initialize config: %w", err)
	}

	return nil
}
