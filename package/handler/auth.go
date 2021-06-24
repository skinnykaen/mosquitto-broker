package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/mqtt-broker"
	"github.com/skinnykaen/mqtt-broker/utils"
	"net/http"
)

func(h *Handler) signUp(c *gin.Context){
	var input mqtt.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	inputPass := input.UserData.Password

	err := h.services.Authorization.CreateUser(input)

	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(c, http.StatusOK, resp)
		return
	}

	h.services.Mosquitto.MosquittoPasswd(input.UserData.Email, inputPass)
	h.services.Mosquitto.MosquittoAcl(input.UserData.Email)

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
		resp := utils.Message(false, err.Error())
		utils.Respond(c, 200, resp)
		return
	}

	resp := utils.Message(true, "Logged in!")
	resp["token"] = token
	utils.Respond(c, http.StatusOK, resp)
}

func(h *Handler) signOut(c *gin.Context){
	h.services.Mosquitto.MosquittoStop()
	fmt.Println("Mosquitto выключено")
	resp := utils.Message(true, "Logout")
	utils.Respond(c, http.StatusOK, resp)
}
