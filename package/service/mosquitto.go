package service

import (
	"github.com/skinnykaen/mqtt-broker/mosquitto"
	"github.com/skinnykaen/mqtt-broker/package/repository"
	"os"
)

type MosquittoService struct {
	repo repository.Mosquitto
}

func NewMosquittoService(repo *repository.Repository) *MosquittoService {
	return &MosquittoService{repo: repo}
}

func (s *MosquittoService) MosquittoRun() {
	args := []string{"-c", os.Getenv("MOSQUITTO_DIR_FILE") + "mosquitto.conf", "-v"}
	go mosquitto.RunCommand(os.Getenv("MOSQUITTO_DIR_EXE") + "mosquitto.exe", args...)
}

func (s *MosquittoService) MosquittoStop() {
	args := []string{"/IM","mosquitto.exe", "/F"}
	go mosquitto.RunCommand("taskkill", args...)
}

func (s *MosquittoService) MosquittoPasswd(email, password string) {
	args := []string{"-b", os.Getenv("MOSQUITTO_DIR_FILE") + "passwd.txt", email, password}
	go mosquitto.RunCommand(os.Getenv("MOSQUITTO_DIR_EXE") + "mosquitto_passwd", args...)
}

func (s *MosquittoService) MosquittoAcl(email string){
	go mosquitto.WriteToAclFile(email) //Записать в acl
}

func (s *MosquittoService) SetMosquittoOn(id uint) (error) {
	return s.repo.SetMosquittoOn(id)
}

func (s *MosquittoService) SetMosquittoOff(id uint) (error) {
	return s.repo.SetMosquittoOff(id)
}

func (s *MosquittoService) GetMosquittoOn(id uint) (bool, error) {
	return s.repo.GetMosquittoOn(id)
}