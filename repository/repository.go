package repository

import (
	"eth-caching-proxy/repository/cloudflare"
)

// Repository provides methods to get blocks from ethereum network
type Repository struct {
	BlockRepository BlockRepository
}

// New creates a repository
func New() *Repository {
	client := cloudflare.GetClient()

	return &Repository{
		BlockRepository: cloudflare.NewBlockRepository(client),
	}
}
