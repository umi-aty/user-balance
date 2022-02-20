package main

import (
	"userbalance/server"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	server.RegisterRoute(r)
}
