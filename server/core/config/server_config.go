package config

import (
	"fmt"

	gpaasgorm "github.com/jackyuan2010/gpaas/server/core/gorm"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	DbType   string             `mapstructure:"dbtype" json:"dbtype" yaml:"dbtype"`
	DbConfig gpaasgorm.DbConfig `mapstructure:"dbconfig" json:"dbconfig" yaml:"dbconfig"`
}

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file %s", err))
	}

	var serverconfig ServerConfig
	if err := v.Unmarshal(&serverconfig); err != nil {
		fmt.Println(err)
	}

	return v
}
