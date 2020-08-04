package main

import (
	"github.com/spf13/viper"
)

// CMCConfig all config needed for interacting with CoinMarket Cap
type CMCConfig struct {
	CmcAPI     string `mapstructure:"CMC_API"`
	CmcBaseURL string `mapstructure:"CMC_BASE_URL"`
}

// getCMCConfig return the configuration in the json file
func getCMCConfig() (CMCConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	// Read Config file
	err := viper.ReadInConfig()
	if err != nil {
		return CMCConfig{}, err
	}

	// Unmarshal the JSON file into a Config structure
	var conf CMCConfig
	err = viper.Unmarshal(&conf)
	if err != nil {
		return CMCConfig{}, err
	}

	return conf, nil
}
