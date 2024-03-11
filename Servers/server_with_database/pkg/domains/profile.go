package domains

import (
	"server_with_database/pkg/models"
	"server_with_database/pkg/types"

	"github.com/gin-gonic/gin"
)

type ProfileRepository interface {
	CreateProfile(profile *models.Profile) error
	GetProfile(profileId int) (models.Profile, error)
}

type ProfileService interface {
	CreateProfile(profile *types.NewProfileData) error
	GetProfile(profileId uint) (types.ProfileInfo, error)
}

type ProfileController interface {
	CreateProfile(g *gin.Context)
	GetProfile(g *gin.Context)
	CreateInRedis(g *gin.Context)
	GetFromRedis(g *gin.Context)
}
