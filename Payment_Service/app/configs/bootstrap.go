package configs

import (
	"github.com/spf13/viper"
)

func Bootstrap() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
