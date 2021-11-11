package repository

import (
	"eth-caching-proxy/repository/cloudflare"
)

type Repository struct {
	BlockRepository BlockRepository
}

func New() *Repository {
	client := cloudflare.GetClient()

	return &Repository{
		BlockRepository: cloudflare.NewBlockRepository(client),
	}
}
