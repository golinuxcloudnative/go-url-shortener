package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Redis  RedisCfg
	Server ServerCfg
}

type RedisCfg struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Port     string `mapstructure:"REDIS_PORT"`
	Password string `mapstructure:"REDIS_PASSWORD"`
}

type ServerCfg struct {
	Port string `mapstructure:"SERVER_PORT"`
}

func NewConfig() (config *Config, err error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config file does not exist:", err)
		}
		return
	}

	redisCfg := &RedisCfg{}

	err = viper.Unmarshal(&redisCfg)
	if err != nil {
		log.Fatalf("unable to unmarshal redis config %v", err)
		return
	}

	serverCfg := &ServerCfg{}

	err = viper.Unmarshal(&serverCfg)
	if err != nil {
		log.Fatalf("unable to unmarshal server config %v", err)
		return
	}

	config = &Config{
		Redis:  *redisCfg,
		Server: *serverCfg,
	}

	return
}
