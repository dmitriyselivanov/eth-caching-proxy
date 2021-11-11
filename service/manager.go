package service

import (
	"errors"
	"eth-caching-proxy/cache"
	"eth-caching-proxy/repository"
)

type Manager struct {
	BlockService EthBlockService
}

func New(repository *repository.Repository, cache *cache.Cache) (*Manager, error) {
	if repository == nil {
		return nil, errors.New("nil repository provided")
	}

	return &Manager{
		BlockService: NewEthBlockService(repository, cache),
	}, nil
}
