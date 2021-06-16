package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/mqtt-broker"
	"github.com/skinnykaen/mqtt-broker/mosquitto"
	"github.com/skinnykaen/mqtt-broker/utils"
	"net/http"
	"os"
)

func(h *Handler) signUp(c *gin.Context){
	var input mqtt.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	inputPass := input.UserData.Password

	err := h.services.Authorization.CreateUser(input)
	fmt.Println(err)
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(c, http.StatusInternalServerError, resp)
		return
	}

	args := []string{"-b", os.Getenv("MOSQUITTO_DIR_FILE") + "passwd", input.UserData.Email, inputPass}
	go mosquitto.RunCommand(os.Getenv("MOSQUITTO_DIR_EXE") + "mosquitto_passwd", args...) //Записать в passwd
	go mosquitto.WriteToAclFile(input.UserData.Email) //Записать в acl

	resp := utils.Message(true, "User was created")
	utils.Respond(c, http.StatusOK, resp)
}

func(h *Handler) signIn(c *gin.Context){
	var input mqtt.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.UserData.Email, input.UserData.Password)
	if err != nil {
		resp := utils.Message(false, "Invalid login Credentials")
		utils.Respond(c, http.StatusInternalServerError, resp)
		return
	}

	resp := utils.Message(true, "Logged in!")
	resp["token"] = token
	utils.Respond(c, http.StatusOK, resp)
}