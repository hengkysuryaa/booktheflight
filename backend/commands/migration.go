package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hengkysuryaa/booktheflight/backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SeatsItineraryParts struct {
	SegmentSeatMaps []SegmentSeatMaps `json:"segmentSeatMaps"`
}

type SegmentSeatMaps struct {
	PassengerSeatMaps []PassengerSeatMap `json:"passengerSeatMaps"`
	Segment           `json:"segment"`
}

type SeatMap struct {
	RowsDisabledCauses []RowDisabledCause `json:"rowsDisabledCauses"`
	Aircraft           string             `json:"aircraft"`
	Cabins             []Cabin            `json:"cabins"`
}

type RowDisabledCause struct {
	RowNumber int    `json:"rowNumber"`
	Cause     string `json:"cause"`
}

type Seat struct {
	SlotCharacteristics    []string     `json:"slotCharacteristics"`
	StorefrontSlotCode     string       `json:"storefrontSlotCode"`
	Available              bool         `json:"available"`
	Entitled               bool         `json:"entitled"`
	FeeWaived              bool         `json:"feeWaived"`
	FreeOfCharge           bool         `json:"freeOfCharge"`
	OriginallySelected     bool         `json:"originallySelected"`
	Code                   string       `json:"code"`
	Designations           []string     `json:"designations"`
	EntitledRuleID         string       `json:"entitledRuleId"`
	FeeWaivedRuleID        string       `json:"feeWaivedRuleId"`
	SeatCharacteristics    []string     `json:"seatCharacteristics"`
	Limitations            []string     `json:"limitations"`
	RefundIndicator        string       `json:"refundIndicator"`
	Prices                 Alternatives `json:"prices"`
	Taxes                  Alternatives `json:"taxes"`
	Total                  Alternatives `json:"total"`
	RawSeatCharacteristics []string     `json:"rawSeatCharacteristics"`
}

type Alternatives struct {
	Alternatives [][]Alternative `json:"alternatives"`
}

type Alternative struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type SeatRow struct {
	RowNumber int      `json:"rowNumber"`
	SeatCodes []string `json:"seatCodes"`
	Seats     []Seat   `json:"seats"`
}

type Cabin struct {
	Deck        string    `json:"deck"`
	SeatColumns []string  `json:"seatColumns"`
	SeatRows    []SeatRow `json:"seatRows"`
	FirstRow    int       `json:"firstRow"`
	LastRow     int       `json:"lastRow"`
}

type Passenger struct {
	PassengerIndex      int    `json:"passengerIndex"`
	PassengerNameNumber string `json:"passengerNameNumber"`
	PassengerDetails    `json:"passengerDetails"`
	PassengerInfo       `json:"passengerInfo"`
	Preferences         Preference `json:"preferences"`
	DocumentInfo        `json:"documentInfo"`
}

type PassengerDetails struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type PassengerInfo struct {
	DateOfBirth string   `json:"dateOfBirth"`
	Gender      string   `json:"gender"`
	Type        string   `json:"type"`
	Emails      []string `json:"emails"`
	Phones      []string `json:"phones"`
	Address     `json:"address"`
}

type Address struct {
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	Postcode    string `json:"postcode"`
	State       string `json:"state"`
	City        string `json:"city"`
	Country     string `json:"country"`
	AddressType string `json:"addressType"`
}

type DocumentInfo struct {
	IssuingCountry string `json:"issuingCountry"`
	CountryOfBirth string `json:"countryOfBirth"`
	DocumentType   string `json:"documentType"`
	Nationality    string `json:"nationality"`
}

type Preference struct {
	SpecialPreferences SpecialPreference `json:"specialPreferences"`
	FrequentFlyer      []FrequentFlyer   `json:"frequentFlyer"`
}

type FrequentFlyer struct {
	Airline    string `json:"airline"`
	Number     string `json:"number"`
	TierNumber int    `json:"tierNumber"`
}

type SpecialPreference struct {
	MealPreference               string   `json:"mealPreference"`
	SeatPreference               string   `json:"seatPreference"`
	SpecialRequests              []string `json:"specialRequests"`
	SpecialServiceRequestRemarks []string `json:"specialServiceRequestRemarks"`
}

type PassengerSeatMap struct {
	SeatSelectionEnabledForPax bool      `json:"seatSelectionEnabledForPax"`
	SeatMap                    SeatMap   `json:"seatMap"`
	Passenger                  Passenger `json:"passenger"`
}

type Segment struct {
	Type                        string `json:"@type"`
	SegmentOfferInformation     `json:"segmentOfferInformation"`
	Duration                    int    `json:"duration"`
	CabinClass                  string `json:"cabinClass"`
	Equipment                   string `json:"equipment"`
	Flight                      `json:"flight"`
	Origin                      string `json:"origin"`
	Destination                 string `json:"destination"`
	Departure                   string `json:"departure"`
	Arrival                     string `json:"arrival"`
	BookingClass                string `json:"bookingClass"`
	LayoverDuration             int    `json:"layoverDuration"`
	FareBasis                   string `json:"fareBasis"`
	SubjectToGovernmentApproval bool   `json:"subjectToGovernmentApproval"`
	SegmentRef                  string `json:"segmentRef"`
}

type SegmentOfferInformation struct {
	FlightsMiles int  `json:"flightsMiles"`
	AwardFare    bool `json:"awardFare"`
}

type Flight struct {
	FlightNumber          int      `json:"flightNumber"`
	OperatingFlightNumber int      `json:"operatingFlightNumber"`
	AirlineCode           string   `json:"airlineCode"`
	OperatingAirlineCode  string   `json:"operatingAirlineCode"`
	StopAirports          []string `json:"stopAirports"`
	DepartureTerminal     string   `json:"departureTerminal"`
	ArrivalTerminal       string   `json:"arrivalTerminal"`
}

type GetSeat struct {
	SeatsItineraryParts []SeatsItineraryParts `json:"seatsItineraryParts"`
	SelectedSeats       []any                 `json:"selectedSeats"`
}

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
		&models.FlightSegment{}, &models.Booking{}, &models.PassengerAddress{}, &models.PassengerDocument{},
	)

	file, err := os.ReadFile("SeatMapResponses.json")
	if err != nil {
		panic("failed to read JSON file")
	}

	// test migration uuid
	flightID, _ := uuid.Parse("04104ded-8380-4d88-9798-0f28e32a616b")
	passengerID, _ := uuid.Parse("3b1ea360-3f82-4f59-918e-b7280d64eb76")

	var response GetSeat
	err = json.Unmarshal(file, &response)
	if err != nil {
		panic("failed to unmarshal JSON")
	}

	for _, itinerary := range response.SeatsItineraryParts {
		for _, segment := range itinerary.SegmentSeatMaps {
			seg := segment.Segment
			flight := models.FlightSegment{
				UUID:                  flightID,
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
					ID:          uint(psm.Passenger.PassengerIndex),
					UUID:        passengerID,
					NameNumber:  psm.Passenger.PassengerNameNumber,
					FirstName:   psm.Passenger.PassengerDetails.FirstName,
					LastName:    psm.Passenger.PassengerDetails.LastName,
					Gender:      psm.Passenger.PassengerInfo.Gender,
					DateOfBirth: parseDate(psm.Passenger.PassengerInfo.DateOfBirth, time.DateOnly),
					Type:        psm.Passenger.PassengerInfo.Type,
					Emails:      psm.Passenger.PassengerInfo.Emails,
					Phones:      psm.Passenger.PassengerInfo.Phones,

					SeatSelectionEnabled: psm.SeatSelectionEnabledForPax,
				}
				db.Create(&passenger)

				doc := models.PassengerDocument{
					PassengerID:    passenger.ID,
					Nationality:    psm.Passenger.DocumentInfo.Nationality,
					DocumentType:   psm.Passenger.DocumentInfo.DocumentType,
					IssuingCountry: psm.Passenger.DocumentInfo.IssuingCountry,
					CountryOfBirth: psm.Passenger.DocumentInfo.CountryOfBirth,
				}
				db.Create(&doc)

				address := models.PassengerAddress{
					PassengerID: passenger.ID,
					Street1:     psm.Passenger.PassengerInfo.Address.Street1,
					Street2:     psm.Passenger.PassengerInfo.Address.Street2,
					City:        psm.Passenger.PassengerInfo.Address.City,
					State:       psm.Passenger.PassengerInfo.Address.State,
					Postcode:    psm.Passenger.PassengerInfo.Address.Postcode,
					AddressType: psm.Passenger.PassengerInfo.Address.AddressType,
					Country:     psm.Passenger.PassengerInfo.Address.Country,
				}
				db.Create(&address)

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
								SeatRowID:           rowModel.ID,
								Code:                seat.Code,
								ColumnLabel:         getSeatColumn(seat.Code),
								SlotCode:            seat.StorefrontSlotCode,
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

							if len(seatModel.Characteristics) > 0 && !seat.Available {
								db.Create(&models.Booking{
									PassengerID:     passenger.ID,
									SeatID:          seatModel.ID,
									FlightSegmentID: flight.ID,
									BookedAt:        time.Now(),
								})
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
