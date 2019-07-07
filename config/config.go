package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Schema struct {
	Database struct {
		Address  string `mapstructure:"address"`
		Database string `mapstructure:"database"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Debug    bool   `mapstructure:"debug"`
		Port     int    `mapstructure:"port"`
	} `mapstructure:"database"`
	API struct {
		Token string `mapstructure:"token"`
	} `mapstructure:"api"`
}

var (
	Config *Schema
)

func init() {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath(".")
	config.AddConfigPath("config/")
	config.AddConfigPath("../config/")
	config.AddConfigPath("../")
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	config.AutomaticEnv()

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}
	err = config.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}
}
