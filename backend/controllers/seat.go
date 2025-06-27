package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hengkysuryaa/booktheflight/backend/responses"
)

type seatController struct {
}

func NewSeat() seatController {
	return seatController{}
}

func (s *seatController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, responses.GetSeat{
		SeatsItineraryParts: responses.SeatsItineraryParts{
			SegmentSeatMaps: responses.SegmentSeatMaps{
				PassengerSeatMaps: []responses.PassengerSeatMap{
					{
						SeatMap:   responses.SeatMap{},
						Passenger: responses.Passenger{},
					},
				},
				Segment: responses.Segment{},
			},
		},
		SelectedSeats: []any{},
	})
}
