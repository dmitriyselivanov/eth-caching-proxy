package service

import (
	"errors"
	"eth-caching-proxy/cache"
	"eth-caching-proxy/repository"
)

// Manager holds references to services
type Manager struct {
	BlockService EthBlockService
}

// New creates new service manager
func New(repository *repository.Repository, cache *cache.Cache) (*Manager, error) {
	if repository == nil {
		return nil, errors.New("nil repository provided")
	}

	return &Manager{
		BlockService: NewEthBlockService(repository, cache),
	}, nil
}
