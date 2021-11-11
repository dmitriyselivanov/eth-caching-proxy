package controller

import (
	"eth-caching-proxy/cache"
	"eth-caching-proxy/repository"
	"eth-caching-proxy/service"
	"github.com/gin-gonic/gin"
	"log"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	repo := repository.New()
	c := cache.New()

	serviceManager, err := service.New(repo, c)
	if err != nil {
		log.Fatalf("error creating service manager")
	}

	blockCtrl := NewEthBlockController(serviceManager)
	blockCtrl.RegisterRoutes(router)

	return router
}
