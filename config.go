package app

import (
	"github.com/spf13/viper"
	"log"
)

func LoadConfig(path string, defaultFunc func(v *viper.Viper)) (v *viper.Viper) {
	v = viper.New()
	defaultFunc(v)
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("load config file err.")
	}
	return
}
