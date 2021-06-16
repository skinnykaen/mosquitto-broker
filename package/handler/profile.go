package handler

import (
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

	resp := utils.Message(true, "getProfile status ok")
	resp["profile"] = user
	utils.Respond(c, http.StatusOK, resp)
}

func (h *Handler) mosquittoOn(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	err = h.services.Profile.SetMosquittoOn(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	resp := utils.Message(true, "Mosquitto On")
	utils.Respond(c, http.StatusOK, resp)
}

func (h *Handler) mosquittoOff (c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	err = h.services.Profile.SetMosquittoOff(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	resp := utils.Message(true, "Mosquitto Off")
	utils.Respond(c, http.StatusOK, resp)
}