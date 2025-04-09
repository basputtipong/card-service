package main

import (
	"card-service/httpserv"
	"card-service/infrastructure"
)

func init() {
	infrastructure.InitConfig()
}

func main() {
	infrastructure.InitMiddleware()
	infrastructure.InitDB()
	httpserv.Run()
}
