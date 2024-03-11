package repositories

import (
	"server_with_database/pkg/domains"
	"server_with_database/pkg/models"

	"gorm.io/gorm"
)

type profileRepo struct {
	db *gorm.DB
}

func SetProfileRepo(DB *gorm.DB) domains.ProfileRepository {
	return profileRepo{
		db: DB,
	}
}

// CreateProfile implements domains.ProfileRepository.
func (pRepo profileRepo) CreateProfile(profile *models.Profile) error {
	if err := pRepo.db.Create(&profile).Error; err != nil {
		return err
	}
	return nil
}

// GetProfile implements domains.ProfileRepository.
func (pRepo profileRepo) GetProfile(profileId int) (models.Profile, error) {
	var existingProfile *models.Profile
	if err := pRepo.db.Where("id = ?", profileId).First(&existingProfile).Error; err != nil {
		return models.Profile{}, err
	}

	return *existingProfile, nil
}
