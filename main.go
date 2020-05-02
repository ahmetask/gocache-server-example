package main

import (
	"fmt"
	"github.com/ahmetask/gocache/v2"
	"os"
	"time"
)

func main() {
	cache := gocache.NewCache(5*time.Minute, 5*time.Second)

	serverConfig := gocache.ServerConfig{
		CachePtr: cache,
		Port:     "10001",
	}

	err := gocache.NewCacheServer(serverConfig)
	fmt.Println(err)
	os.Exit(-1)
}
