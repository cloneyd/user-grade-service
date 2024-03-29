package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PublicServer  ServerConfig
	PrivateServer ServerConfig
	StanConn      StanConfig
}

type ServerConfig struct {
	Addr string
}

type StanConfig struct {
	ClusterId string
	ClientId  string
	Subject   string
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath("/")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
