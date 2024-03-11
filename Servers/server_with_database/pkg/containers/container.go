package containers

import (
	"server_with_database/pkg/connection"
	"server_with_database/pkg/controllers"
	"server_with_database/pkg/repositories"
	"server_with_database/pkg/services"

	"github.com/gin-gonic/gin"
)

func InventoryRoutes(g *gin.Engine) {

	db := connection.GetDB()
	// redis := connection.GetRedis()

	profileRepo := repositories.SetProfileRepo(db)
	profileService := services.SetProfileService(profileRepo)
	profileController := controllers.SetProfileController(profileService)

	profile := g.Group("/profile")
	{
		profile.POST("/create", profileController.CreateProfile)
		profile.GET("/get", profileController.GetProfile)
		profile.POST("/redis/create", profileController.CreateInRedis)
		profile.GET("/redis/get", profileController.GetFromRedis)
	}
}
