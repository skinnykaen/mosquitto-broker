package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/mqtt-broker/utils"
	"net/http"
)

func (h *Handler) getProfile(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}
	user, err := h.services.Profile.GetProfile(id)

	if err != nil {
		newErrorResponse(c,  http.StatusOK, err.Error())
		return
	}

	if(user.UserData.MosquittoOn){
		h.services.Mosquitto.MosquittoRun()
		fmt.Println("Mosquitto включено")
	}

	resp := utils.Message(true, "getProfile status ok")
	resp["profile"] = user
	utils.Respond(c, http.StatusOK, resp)
}
