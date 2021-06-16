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

func (s *ProfileService) SetMosquittoOn(id uint) (error) {
	//args := []string{"-c", os.Getenv("MOSQUITTO_DIR_EXE") + "mosquitto.conf"}
	//fmt.Println(args)
	//go mosquitto.RunCommand(os.Getenv("MOSQUITTO_DIR_EXE") + "mosquitto.exe")
	return s.repo.SetMosquittoOn(id)
}

func (s *ProfileService) SetMosquittoOff(id uint) (error) {
	return s.repo.SetMosquittoOff(id)
}