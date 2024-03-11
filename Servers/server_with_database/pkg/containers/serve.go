package containers

import (
	"log"
	"net/http"
	"server_with_database/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Serve(g *gin.Engine) {
	config := utils.GetConfig()
	InventoryRoutes(g)

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: g,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start", err)
	}
}
