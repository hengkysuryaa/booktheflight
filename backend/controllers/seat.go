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
		SeatsItineraryParts: []responses.SeatsItineraryParts{
			{
				SegmentSeatMaps: []responses.SegmentSeatMaps{
					{
						PassengerSeatMaps: []responses.PassengerSeatMap{
							{
								SeatMap: responses.SeatMap{
									RowsDisabledCauses: []responses.RowDisabledCause{},
									Cabins: []responses.Cabin{
										{
											SeatColumns: []string{},
											SeatRows: []responses.SeatRow{
												{
													SeatCodes: []string{},
													Seats: []responses.Seat{
														{
															RawSeatCharacteristics: []string{},
															SeatCharacteristics:    []string{},
															Designations:           []string{},
															Limitations:            []string{},
															Prices: struct {
																Alternatives [][]responses.Alternative "json:\"alternatives\""
															}{
																Alternatives: [][]responses.Alternative{
																	{
																		{
																			Amount:   1,
																			Currency: "MYR",
																		},
																	},
																},
															},
															Taxes: struct {
																Alternatives [][]responses.Alternative "json:\"alternatives\""
															}{
																Alternatives: [][]responses.Alternative{
																	{},
																},
															},
															Total: struct {
																Alternatives [][]responses.Alternative "json:\"alternatives\""
															}{
																Alternatives: [][]responses.Alternative{
																	{},
																},
															},
														},
													},
												},
											},
										},
									},
									Aircraft: "",
								},
								Passenger: responses.Passenger{
									Preferences: responses.Preference{
										SpecialPreferences: responses.SpecialPreference{
											SpecialRequests:              []string{},
											SpecialServiceRequestRemarks: []string{},
										},
										FrequentFlyer: []responses.FrequentFlyer{},
									},
									PassengerDetails: responses.PassengerDetails{},
									PassengerInfo: responses.PassengerInfo{
										Address: responses.Address{},
										Emails:  []string{},
										Phones:  []string{},
									},
								},
							},
						},
						Segment: responses.Segment{
							SegmentOfferInformation: responses.SegmentOfferInformation{},
							Flight: responses.Flight{
								StopAirports: []any{},
							},
						},
					},
				},
			},
		},
		SelectedSeats: []any{},
	})
}
