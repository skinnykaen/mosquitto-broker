package service

import (
	"github.com/skinnykaen/mqtt-broker"
	"github.com/skinnykaen/mqtt-broker/package/repository"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfileService(repo *repository.Repository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) GetProfile(id uint) (mqtt.User, error) {
	return  s.repo.GetProfile(id)
}