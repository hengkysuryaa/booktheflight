package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hengkysuryaa/booktheflight/backend/services"
)

type seatController struct {
	seatSvc services.ISeatService
}

func NewSeat(seatSvc services.ISeatService) seatController {
	return seatController{
		seatSvc: seatSvc,
	}
}

func (s *seatController) Get(c *gin.Context) {
	flightID := c.Request.URL.Query().Get("flight_id")
	fUuid, err := uuid.Parse(flightID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	passengerID := c.Request.URL.Query().Get("passenger_id")
	pUuid, err := uuid.Parse(passengerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := s.seatSvc.GetSeats(c.Request.Context(), fUuid, pUuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
