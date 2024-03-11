package services

import (
	"server_with_database/pkg/domains"
	"server_with_database/pkg/models"
	"server_with_database/pkg/types"
)

type profileService struct {
	pRepo domains.ProfileRepository
}

func SetProfileService(profileRepository domains.ProfileRepository) domains.ProfileService {
	return profileService{
		pRepo: profileRepository,
	}
}

// CreateProfile implements domains.ProfileService.
func (pService profileService) CreateProfile(profile *types.NewProfileData) error {
	newProfile := &models.Profile{
		Username: profile.Username,
		Email:    profile.Email,
		Password: profile.Password,
	}
	if err := pService.pRepo.CreateProfile(newProfile); err != nil {
		return err
	}
	return nil
}

// GetProfile implements domains.ProfileService.
func (pService profileService) GetProfile(profileId uint) (types.ProfileInfo, error) {
	existingProfile, err := pService.pRepo.GetProfile(int(profileId))
	if err != nil {
		return types.ProfileInfo{}, err
	}
	return types.ProfileInfo{
		ID:       existingProfile.ID,
		Username: existingProfile.Username,
		Email:    existingProfile.Email,
		Password: existingProfile.Password,
	}, nil
}
