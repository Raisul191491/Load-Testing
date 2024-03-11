package main

import (
	"server_with_database/pkg/containers"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	containers.Serve(g)
}
