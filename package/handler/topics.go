package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/mqtt-broker"
	"github.com/skinnykaen/mqtt-broker/utils"
	"net/http"
	"strconv"
)

func(h *Handler) createTopic(c *gin.Context){
	var input mqtt.Topic
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	idUser, err := getUserId(c)
	if err != nil {
		return
	}
	input.Id_User = idUser

	idTopic, err := h.services.Topics.CreateTopic(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	input.Id = idTopic
	resp := utils.Message(true, "Topic was created")
	resp["topic"] = input
	utils.Respond(c, http.StatusOK, resp)
}

func(h *Handler) getAllTopics(c *gin.Context){
	id, err := getUserId(c)
	if err != nil {
		return
	}

	topics, err := h.services.Topics.GetTopics(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp := utils.Message(true, "Topics have been received")
	resp["topics"] = topics
	utils.Respond(c, http.StatusOK, resp)
}

func(h *Handler) deleteTopic(c *gin.Context){
	idTopic, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	fmt.Println(idTopic)
	err = h.services.Topics.Delete(idTopic)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	resp := utils.Message(true, "Topic has been deleted")
	utils.Respond(c, http.StatusOK, resp)
}