package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hengkysuryaa/booktheflight/backend/models"
	"github.com/hengkysuryaa/booktheflight/backend/responses"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migration() {
	godotenv.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("failed to connect to database: %v", err.Error())
	}

	db.AutoMigrate(
		&models.Aircraft{}, &models.Cabin{}, &models.SeatRow{}, &models.Seat{},
		&models.Price{}, &models.SeatTax{}, &models.SeatTotal{}, &models.RowDisabledCause{},
		&models.Passenger{}, &models.PassengerPreference{}, &models.FrequentFlyer{},
		&models.FlightSegment{}, &models.Booking{},
	)

	file, err := os.ReadFile("SeatMapResponses.json")
	if err != nil {
		panic("failed to read JSON file")
	}

	var response responses.GetSeat
	err = json.Unmarshal(file, &response)
	if err != nil {
		panic("failed to unmarshal JSON")
	}

	for _, itinerary := range response.SeatsItineraryParts {
		for _, segment := range itinerary.SegmentSeatMaps {
			seg := segment.Segment
			flight := models.FlightSegment{
				UUID:                  uuid.New(),
				Origin:                seg.Origin,
				Destination:           seg.Destination,
				Departure:             parseDate(seg.Departure, "2006-01-02T15:04:05"),
				Arrival:               parseDate(seg.Arrival, "2006-01-02T15:04:05"),
				FlightNumber:          seg.FlightNumber,
				AirlineCode:           seg.AirlineCode,
				Equipment:             seg.Equipment,
				CabinClass:            seg.CabinClass,
				BookingClass:          seg.BookingClass,
				FareBasis:             seg.FareBasis,
				FlightsMiles:          seg.FlightsMiles,
				AwardFare:             seg.AwardFare,
				Duration:              seg.Duration,
				LayoverDuration:       seg.LayoverDuration,
				SubjectToGovtApproval: seg.SubjectToGovernmentApproval,
				SegmentRef:            seg.SegmentRef,
				OperatingFlightNumber: seg.OperatingFlightNumber,
				OperatingAirlineCode:  seg.OperatingAirlineCode,
				StopAirports:          seg.StopAirports,
				DepartureTerminal:     seg.DepartureTerminal,
				ArrivalTerminal:       seg.ArrivalTerminal,
			}
			db.Create(&flight)

			for _, psm := range segment.PassengerSeatMaps {
				passenger := models.Passenger{
					UUID:                 uuid.New(),
					FirstName:            psm.Passenger.PassengerDetails.FirstName,
					LastName:             psm.Passenger.PassengerDetails.LastName,
					Gender:               psm.Passenger.PassengerInfo.Gender,
					DateOfBirth:          parseDate(psm.Passenger.PassengerInfo.DateOfBirth, time.DateOnly),
					Type:                 psm.Passenger.PassengerInfo.Type,
					Emails:               psm.Passenger.PassengerInfo.Emails,
					Phones:               psm.Passenger.PassengerInfo.Phones,
					Nationality:          psm.Passenger.DocumentInfo.Nationality,
					DocumentType:         psm.Passenger.DocumentInfo.DocumentType,
					IssuingCountry:       psm.Passenger.DocumentInfo.IssuingCountry,
					CountryOfBirth:       psm.Passenger.DocumentInfo.CountryOfBirth,
					Street1:              psm.Passenger.PassengerInfo.Address.Street1,
					Street2:              psm.Passenger.PassengerInfo.Address.Street2,
					City:                 psm.Passenger.PassengerInfo.Address.City,
					State:                psm.Passenger.PassengerInfo.Address.State,
					Postcode:             psm.Passenger.PassengerInfo.Address.Postcode,
					AddressType:          psm.Passenger.PassengerInfo.Address.AddressType,
					SeatSelectionEnabled: psm.SeatSelectionEnabledForPax,
				}
				db.Create(&passenger)

				pp := models.PassengerPreference{
					PassengerID:           passenger.ID,
					MealPreference:        psm.Passenger.Preferences.SpecialPreferences.MealPreference,
					SeatPreference:        psm.Passenger.Preferences.SpecialPreferences.SeatPreference,
					SpecialRequests:       psm.Passenger.Preferences.SpecialPreferences.SpecialRequests,
					ServiceRequestRemarks: psm.Passenger.Preferences.SpecialPreferences.SpecialServiceRequestRemarks,
				}
				db.Create(&pp)

				for _, ff := range psm.Passenger.Preferences.FrequentFlyer {
					db.Create(&models.FrequentFlyer{
						PassengerID: passenger.ID,
						Airline:     ff.Airline,
						Number:      ff.Number,
						TierNumber:  ff.TierNumber,
					})
				}

				aircraft := models.Aircraft{Code: psm.SeatMap.Aircraft}
				db.Create(&aircraft)

				for _, cabin := range psm.SeatMap.Cabins {
					cabinModel := models.Cabin{
						AircraftID:  aircraft.ID,
						Deck:        cabin.Deck,
						SeatColumns: cabin.SeatColumns,
						FirstRow:    cabin.FirstRow,
						LastRow:     cabin.LastRow,
					}
					db.Create(&cabinModel)

					for _, rc := range psm.SeatMap.RowsDisabledCauses {
						cause := models.RowDisabledCause{
							CabinID:   cabinModel.ID,
							RowNumber: rc.RowNumber,
							Cause:     rc.Cause,
						}
						db.Create(&cause)
					}

					for _, row := range cabin.SeatRows {
						rowModel := models.SeatRow{
							CabinID:   cabinModel.ID,
							RowNumber: row.RowNumber,
							Codes:     row.SeatCodes,
						}
						db.Create(&rowModel)

						for _, seat := range row.Seats {
							seatModel := models.Seat{
								SeatRowID:   rowModel.ID,
								Code:        seat.Code,
								ColumnLabel: getSeatColumn(seat.Code),
								SlotCode:    seat.StorefrontSlotCode,
								//Available:          seat.Available,
								Entitled:            seat.Entitled,
								FeeWaived:           seat.FeeWaived,
								FreeOfCharge:        seat.FreeOfCharge,
								OriginallySelected:  seat.OriginallySelected,
								RefundIndicator:     seat.RefundIndicator,
								Segment:             flight.SegmentRef,
								Characteristics:     seat.SeatCharacteristics,
								RawCharacteristics:  seat.RawSeatCharacteristics,
								Limitations:         seat.Limitations,
								Designations:        seat.Designations,
								EntitledRuleId:      seat.EntitledRuleID,
								FeeWaivedRuleId:     seat.FeeWaivedRuleID,
								SlotCharacteristics: seat.SlotCharacteristics,
							}
							db.Create(&seatModel)

							for _, group := range seat.Prices.Alternatives {
								for _, price := range group {
									db.Create(&models.Price{SeatID: seatModel.ID, Currency: price.Currency, Amount: price.Amount})
								}
							}

							for _, group := range seat.Taxes.Alternatives {
								for _, tax := range group {
									db.Create(&models.SeatTax{SeatID: seatModel.ID, Currency: tax.Currency, Amount: tax.Amount})
								}
							}

							for _, group := range seat.Total.Alternatives {
								for _, total := range group {
									db.Create(&models.SeatTotal{SeatID: seatModel.ID, Currency: total.Currency, Amount: total.Amount})
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("complete")
}

func parseDate(s string, format string) time.Time {
	t, _ := time.Parse(format, s)
	return t
}

func getSeatColumn(code string) string {
	if len(code) > 0 {
		return string(code[len(code)-1])
	}
	return ""
}
