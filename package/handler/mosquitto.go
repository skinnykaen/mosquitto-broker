package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/mqtt-broker"
	"github.com/skinnykaen/mqtt-broker/utils"
	"net/http"
)

func (h *Handler) MosquittoLaunch(c *gin.Context) {

	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input mqtt.Mosquitto
	//Получаем данные
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var resp map[string] interface{}

	if (input.MosquittoOn){
		err = h.services.Mosquitto.SetMosquittoOn(id)

		h.services.MosquittoRun()

		resp = utils.Message(true, "Mosquitto On")
		fmt.Println("Mosquitto включено")
	}else{
		err = h.services.Mosquitto.SetMosquittoOff(id)

		h.services.MosquittoStop()

		resp = utils.Message(true, "Mosquitto Off")
		fmt.Println("Mosquitto остановлено")
	}

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	utils.Respond(c, http.StatusOK, resp)
}
