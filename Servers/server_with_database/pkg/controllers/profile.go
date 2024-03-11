package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server_with_database/pkg/connection"
	"server_with_database/pkg/consts"
	"server_with_database/pkg/domains"
	"server_with_database/pkg/types"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var cnt int = 1
var mapID int = 1

type localMap struct {
	mp map[int]types.ProfileInfo
}

type profileController struct {
	pService domains.ProfileService
}

func SetProfileController(profileService domains.ProfileService) domains.ProfileController {
	return profileController{
		pService: profileService,
	}
}

// CreateProfile implements domains.ProfileController.
func (pController profileController) CreateProfile(g *gin.Context) {
	input := &types.NewProfileData{}
	if err := g.ShouldBindWith(input, binding.JSON); err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, "Invalid input")
		return
	}

	if err := pController.pService.CreateProfile(input); err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError, "Profile creation failed")
	}

	g.JSON(http.StatusCreated, "Profile created successfully")
}

// GetProfile implements domains.ProfileController.
func (pController profileController) GetProfile(g *gin.Context) {
	tempID := g.Query(consts.ProfileID)
	fmt.Println(tempID)
	profileID, err := strconv.Atoi(tempID)
	if err != nil {
		g.JSON(http.StatusExpectationFailed, "Couldn't parse profile ID")
	}
	existingProfile, err := pController.pService.GetProfile(uint(profileID))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError, "Profile was not found")
	}

	g.JSON(http.StatusAccepted, existingProfile)
}

func (pController profileController) CreateInRedis(g *gin.Context) {
	redis := connection.GetRedis()
	input := &types.NewProfileData{}
	if err := g.ShouldBindWith(input, binding.JSON); err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, "Invalid input")
		return
	}
	profileJSON, err := json.Marshal(input)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize profile"})
		return
	}

	err = redis.Set(g, fmt.Sprintf("%d", cnt), profileJSON, 0).Err()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store profile in Redis"})
		return
	}
	cnt++

	g.JSON(http.StatusCreated, gin.H{"message": "Profile created successfully"})

}

func (pController profileController) GetFromRedis(g *gin.Context) {
	redis := connection.GetRedis()
	profileID := g.Query(consts.ProfileID)
	fmt.Println(profileID)
	val, err := redis.Get(g, profileID).Result()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve profile from Redis"})
		return
	}
	g.JSON(http.StatusOK, val)

}

// func (pController profileController) WriteInLocalMap(g *gin.Context) {
// 	input := &types.NewProfileData{}
// 	if err := g.ShouldBindWith(input, binding.JSON); err != nil {
// 		g.AbortWithStatusJSON(http.StatusBadRequest, "Invalid input")
// 		return
// 	}

// 	localMap.
// }
