package config

import (
	"github.com/spf13/viper"
	"log"
	"path"
	"runtime"
	"sync"
)

// Server represents server configuration
type Server struct {
	RunMode     string `mapstructure:"runmode"`
	HTTPPort    string `mapstructure:"httpport"`
	ReadTimeout int    `mapstructure:"readtimeout"`
	WriteTimeout int    `mapstructure:"writetimeout"`
}

// Cache represents cache configuration
type Cache struct {
	MaxBlocks int `mapstructure:"maxblocks"`
}

// Cloudflare represents cloudflare configuration
type Cloudflare struct {
	URL string
}

// Config represents configuration with server, cache and cloudflare configs
type Config struct {
	Server     Server     `mapstructure:"server"`
	Cache      Cache      `mapstructure:"cache"`
	Cloudflare Cloudflare `mapstructure:"cloudflare"`
}

var config *Config

// GetConfig returns a config
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
