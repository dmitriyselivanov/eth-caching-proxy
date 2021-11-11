package cloudflare

import (
	"eth-caching-proxy/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"sync"
)

var client *ethclient.Client

// GetClient returns ethclient instance
func GetClient() *ethclient.Client {
	var once sync.Once
	once.Do(func() {
		conf := config.GetConfig()

		c, err := ethclient.Dial(conf.Cloudflare.URL)
		if err != nil {
			log.Fatalf("error connecting to cloudflare")
		}

		client = c
	})

	return client
}
