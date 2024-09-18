package config

import (
	"strings"

	"github.com/spf13/viper"
)

func GetConfig(env string, confFiles map[string]string) (*viper.Viper, error) {
	conf := viper.New()
	conf.SetDefault("environment", env)

	// Conf Env
	conf.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "_", "__")) // APP_DATA__BASE_PASS -> app.data_base.pass
	conf.AutomaticEnv()                                              // Automatically load Env variables

	// Conf Files
	//conf.SetConfigType("yaml") 				// We're using yaml
	conf.SetConfigName(env)                   // Search for a config file that matches our environment
	conf.AddConfigPath("./src/config/" + env) // look for config in the working directory
	conf.ReadInConfig()                       // Find and read the config file

	// Read additional files
	for confFile := range confFiles {
		conf.SetConfigName(confFile)
		conf.MergeInConfig()
	}

	return conf, nil
}
