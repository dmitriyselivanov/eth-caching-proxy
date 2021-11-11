package main

import (
	"eth-caching-proxy/config"
	"eth-caching-proxy/controller"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	conf := config.GetConfig()

	endPoint := fmt.Sprintf(":%s", conf.Server.HTTPPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        controller.NewRouter(),
		ReadTimeout:    time.Duration(conf.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(conf.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Server err: %v", err)
	}
}
