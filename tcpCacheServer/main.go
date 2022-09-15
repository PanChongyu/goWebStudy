package main

import (
	"httpBuffer/HTTP"
	"httpBuffer/cache"
	"tcpCacheServer/tcp"
)

func main() {
	ca := cache.New("inmemory")
	go tcp.New(ca).Listen()
	HTTP.New(ca).Listen()
}
