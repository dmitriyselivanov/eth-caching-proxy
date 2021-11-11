package config

import (
	"github.com/spf13/viper"
	"log"
	"path"
	"runtime"
	"sync"
)

type Server struct {
	RunMode      string `mapstructure:"runmode"`
	HttpPort     string `mapstructure:"httpport"`
	ReadTimeout  int    `mapstructure:"readtimeout"`
	WriteTimeout int    `mapstructure:"writetimeout"`
}

type Cache struct {
	MaxBlocks int `mapstructure:"maxblocks"`
}

type Cloudflare struct {
	Url string
}

type Config struct {
	Server     Server     `mapstructure:"server"`
	Cache      Cache      `mapstructure:"cache"`
	Cloudflare Cloudflare `mapstructure:"cloudflare"`
}

var config *Config

func GetConfig() *Config {
	var once sync.Once
	once.Do(func() {
		config = newConfig()
	})

	return config
}

func newConfig() *Config {
	_, filename, _, _ := runtime.Caller(1)
	dirname := path.Dir(filename)

	viper.SetConfigFile(path.Join(dirname, "./config.yaml"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("fatal error while reading the config file: \n", err)
	}

	var config Config
	if err = viper.Unmarshal(&config); err != nil {
		log.Fatalln("fatal error while unmarshaling the config file: \n", err)
	}

	return &config
}
