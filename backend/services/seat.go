package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hengkysuryaa/booktheflight/backend/models"
	"github.com/hengkysuryaa/booktheflight/backend/repository"
	"github.com/hengkysuryaa/booktheflight/backend/responses"
)

type ISeatService interface {
	GetSeats(ctx context.Context, flightID, passangerID uuid.UUID) (res responses.GetSeat, err error)
}

type seatService struct {
	repo repository.IRepository
}

func NewSeatService(repo repository.IRepository) ISeatService {
	return &seatService{
		repo: repo,
	}
}

func (s *seatService) GetSeats(ctx context.Context, flightID, passengerID uuid.UUID) (responses.GetSeat, error) {
	var res responses.GetSeat

	segment, err := s.repo.GetFlight(ctx, flightID)
	if err != nil {
		return res, err
	}

	aircraft, err := s.repo.GetAircraft(ctx, segment.Equipment)
	if err != nil {
		return res, err
	}

	passenger, err := s.repo.GetPassenger(ctx, passengerID)
	if err != nil {
		return res, err
	}

	cabins := mapCabinsToResponse(aircraft.Cabins, segment.Bookings)
	rowsDisabled := mapRowDisabledToResponse(aircraft.Cabins)
	segmentRes := mapSegmentToResponse(segment)
	passengerRes := mapPassengerToResponse(passenger)

	res = responses.GetSeat{
		SeatsItineraryParts: []responses.SeatsItineraryParts{
			{
				SegmentSeatMaps: []responses.SegmentSeatMaps{
					{
						PassengerSeatMaps: []responses.PassengerSeatMap{
							{
								SeatSelectionEnabledForPax: passenger.SeatSelectionEnabled,
								SeatMap: responses.SeatMap{
									Aircraft:           aircraft.Code,
									Cabins:             cabins,
									RowsDisabledCauses: rowsDisabled,
								},
								Passenger: passengerRes,
							},
						},
						Segment: segmentRes,
					},
				},
			},
		},
		SelectedSeats: []any{},
	}

	return res, nil
}

func mapPassengerToResponse(p models.Passenger) responses.Passenger {
	ff := []responses.FrequentFlyer{}
	for _, f := range p.FrequentFlyers {
		ff = append(ff, responses.FrequentFlyer{
			Airline:    f.Airline,
			Number:     f.Number,
			TierNumber: f.TierNumber,
		})
	}

	return responses.Passenger{
		PassengerIndex:      int(p.ID),
		PassengerNameNumber: p.NameNumber,
		PassengerDetails: responses.PassengerDetails{
			FirstName: p.FirstName,
			LastName:  p.LastName,
		},
		DocumentInfo: responses.DocumentInfo{
			IssuingCountry: p.Document.IssuingCountry,
			CountryOfBirth: p.Document.CountryOfBirth,
			DocumentType:   p.Document.DocumentType,
			Nationality:    p.Document.Nationality,
		},
		PassengerInfo: responses.PassengerInfo{
			DateOfBirth: p.DateOfBirth.Format(time.DateOnly),
			Gender:      p.Gender,
			Type:        p.Type,
			Emails:      p.Emails,
			Phones:      p.Phones,
			Address: responses.Address{
				Street1:     p.Address.Street1,
				Street2:     p.Address.Street2,
				Postcode:    p.Address.Postcode,
				State:       p.Address.State,
				City:        p.Address.City,
				Country:     p.Address.Country,
				AddressType: p.Address.AddressType,
			},
		},
		Preferences: responses.Preference{
			SpecialPreferences: responses.SpecialPreference{
				MealPreference:               p.Preferences.MealPreference,
				SeatPreference:               p.Preferences.SeatPreference,
				SpecialRequests:              p.Preferences.SpecialRequests,
				SpecialServiceRequestRemarks: p.Preferences.ServiceRequestRemarks,
			},
			FrequentFlyer: ff,
		},
	}
}

func mapSegmentToResponse(f models.FlightSegment) responses.Segment {
	return responses.Segment{
		Type: "Segment",
		SegmentOfferInformation: responses.SegmentOfferInformation{
			FlightsMiles: f.FlightsMiles,
			AwardFare:    f.AwardFare,
		},
		Duration:   f.Duration,
		CabinClass: f.CabinClass,
		Equipment:  f.Equipment,
		Flight: responses.Flight{
			FlightNumber:          f.FlightNumber,
			OperatingFlightNumber: f.OperatingFlightNumber,
			AirlineCode:           f.AirlineCode,
			OperatingAirlineCode:  f.OperatingAirlineCode,
			StopAirports:          f.StopAirports,
			DepartureTerminal:     f.DepartureTerminal,
			ArrivalTerminal:       f.ArrivalTerminal,
		},
		Origin:                      f.Origin,
		Destination:                 f.Destination,
		Departure:                   f.Departure.UTC().Format("2006-01-02T15:04:05"),
		Arrival:                     f.Arrival.UTC().Format("2006-01-02T15:04:05"),
		BookingClass:                f.BookingClass,
		LayoverDuration:             f.LayoverDuration,
		FareBasis:                   f.FareBasis,
		SubjectToGovernmentApproval: f.SubjectToGovtApproval,
		SegmentRef:                  f.SegmentRef,
	}
}

func mapCabinsToResponse(cabins []models.Cabin, bookings []models.Booking) []responses.Cabin {
	var out []responses.Cabin
	for _, c := range cabins {
		var seatRows []responses.SeatRow
		for _, r := range c.SeatRows {
			var seats []responses.Seat
			for _, s := range r.Seats {
				seats = append(seats, mapSeatToResponse(s, bookings))
			}
			seatRows = append(seatRows, responses.SeatRow{
				RowNumber: r.RowNumber,
				SeatCodes: r.Codes,
				Seats:     seats,
			})
		}
		out = append(out, responses.Cabin{
			Deck:        c.Deck,
			SeatColumns: c.SeatColumns,
			FirstRow:    c.FirstRow,
			LastRow:     c.LastRow,
			SeatRows:    seatRows,
		})
	}
	return out
}

func mapRowDisabledToResponse(cabins []models.Cabin) []responses.RowDisabledCause {
	var out []responses.RowDisabledCause
	for _, c := range cabins {
		for _, rd := range c.RowDisabledCauses {
			out = append(out, responses.RowDisabledCause{
				RowNumber: rd.RowNumber,
				Cause:     rd.Cause,
			})
		}
	}
	return out
}

func mapSeatToResponse(s models.Seat, bookings []models.Booking) responses.Seat {
	bookedMap := make(map[uint]bool)
	for _, b := range bookings {
		bookedMap[b.SeatID] = true
	}

	if len(s.Characteristics) > 0 {
		return responses.Seat{
			StorefrontSlotCode:  s.SlotCode,
			Entitled:            s.Entitled,
			FeeWaived:           s.FeeWaived,
			FreeOfCharge:        s.FreeOfCharge,
			OriginallySelected:  s.OriginallySelected,
			SlotCharacteristics: s.SlotCharacteristics,
			Available:           !bookedMap[s.ID],
			SeatDetail: &responses.SeatDetail{
				Code:                   s.Code,
				Designations:           s.Designations,
				EntitledRuleID:         s.EntitledRuleId,
				FeeWaivedRuleID:        s.FeeWaivedRuleId,
				SeatCharacteristics:    s.Characteristics,
				Limitations:            s.Limitations,
				RefundIndicator:        s.RefundIndicator,
				RawSeatCharacteristics: s.RawCharacteristics,
				Prices:                 mapAlt(s.Prices),
				Taxes:                  mapAlt(s.Taxes),
				Total:                  mapAlt(s.Totals),
			},
		}
	} else {
		return responses.Seat{
			StorefrontSlotCode:  s.SlotCode,
			Available:           false,
			Entitled:            s.Entitled,
			FeeWaived:           s.FeeWaived,
			FreeOfCharge:        s.FreeOfCharge,
			OriginallySelected:  s.OriginallySelected,
			SlotCharacteristics: s.SlotCharacteristics,
		}
	}
}

func mapAlt[T any](items []T) struct {
	Alternatives [][]responses.Alternative `json:"alternatives"`
} {
	var alts [][]responses.Alternative
	for _, item := range items {
		switch v := any(item).(type) {
		case models.Price:
			alts = append(alts, []responses.Alternative{{Amount: v.Amount, Currency: v.Currency}})
		case models.SeatTax:
			alts = append(alts, []responses.Alternative{{Amount: v.Amount, Currency: v.Currency}})
		case models.SeatTotal:
			alts = append(alts, []responses.Alternative{{Amount: v.Amount, Currency: v.Currency}})
		}
	}
	return struct {
		Alternatives [][]responses.Alternative `json:"alternatives"`
	}{Alternatives: alts}
}
