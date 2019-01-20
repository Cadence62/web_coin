package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("sys")
	//viper.SetConfigName("coin")
	viper.AddConfigPath("C:\\Users\\menfan\\go\\src\\web_coin\\config")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}
}
