package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Aircraft struct {
	ID     uint   `gorm:"primaryKey"`
	Code   string `gorm:"not null;uniqueIndex"`
	Cabins []Cabin
}

type Cabin struct {
	ID                uint `gorm:"primaryKey"`
	AircraftID        uint
	Deck              string
	SeatColumns       pq.StringArray `gorm:"type:text[]"`
	FirstRow          int
	LastRow           int
	SeatRows          []SeatRow
	RowDisabledCauses []RowDisabledCause
}

type SeatRow struct {
	ID        uint `gorm:"primaryKey"`
	CabinID   uint
	RowNumber int
	Codes     pq.StringArray `gorm:"type:text[]"`
	Seats     []Seat
}

type Seat struct {
	ID                  uint `gorm:"primaryKey"`
	SeatRowID           uint
	Code                string
	ColumnLabel         string
	SlotCode            string
	Entitled            bool
	FeeWaived           bool
	EntitledRuleId      string
	FeeWaivedRuleId     string
	FreeOfCharge        bool
	OriginallySelected  bool
	RefundIndicator     string
	Segment             string
	Characteristics     pq.StringArray `gorm:"type:text[]"`
	RawCharacteristics  pq.StringArray `gorm:"type:text[]"`
	SlotCharacteristics pq.StringArray `gorm:"type:text[]"`
	Limitations         pq.StringArray `gorm:"type:text[]"`
	Designations        pq.StringArray `gorm:"type:text[]"`
	Prices              []Price
	Taxes               []SeatTax
	Totals              []SeatTotal
}

type Price struct {
	ID       uint `gorm:"primaryKey"`
	SeatID   uint
	Currency string
	Amount   float64
}

type SeatTax struct {
	ID       uint `gorm:"primaryKey"`
	SeatID   uint
	Amount   float64
	Currency string
}

type SeatTotal struct {
	ID       uint `gorm:"primaryKey"`
	SeatID   uint
	Amount   float64
	Currency string
}

type RowDisabledCause struct {
	ID        uint `gorm:"primaryKey"`
	CabinID   uint
	RowNumber int
	Cause     string
}

type Passenger struct {
	ID                   uint      `gorm:"primaryKey"`
	UUID                 uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();uniqueIndex"`
	NameNumber           string
	FirstName            string
	LastName             string
	Type                 string
	SeatSelectionEnabled bool
	Gender               string
	DateOfBirth          time.Time
	Emails               pq.StringArray `gorm:"type:text[]"`
	Phones               pq.StringArray `gorm:"type:text[]"`
	Document             PassengerDocument
	Address              PassengerAddress
	Preferences          PassengerPreference
	FrequentFlyers       []FrequentFlyer
}

type PassengerAddress struct {
	ID          uint `gorm:"primaryKey"`
	PassengerID uint
	Street1     string
	Street2     string
	City        string
	State       string
	Country     string
	Postcode    string
	AddressType string
}

type PassengerDocument struct {
	ID             uint `gorm:"primaryKey"`
	PassengerID    uint
	Nationality    string
	DocumentType   string
	IssuingCountry string
	CountryOfBirth string
}

type PassengerPreference struct {
	ID                    uint `gorm:"primaryKey"`
	PassengerID           uint
	MealPreference        string
	SeatPreference        string
	SpecialRequests       pq.StringArray `gorm:"type:text[]"`
	ServiceRequestRemarks pq.StringArray `gorm:"type:text[]"`
}

type FrequentFlyer struct {
	ID          uint `gorm:"primaryKey"`
	PassengerID uint
	Airline     string
	Number      string
	TierNumber  int
}

type FlightSegment struct {
	ID                    uint      `gorm:"primaryKey"`
	UUID                  uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();uniqueIndex"`
	Origin                string
	Destination           string
	Departure             time.Time
	Arrival               time.Time
	FlightNumber          int
	AirlineCode           string
	OperatingFlightNumber int
	OperatingAirlineCode  string
	StopAirports          []string `gorm:"type:text[]"`
	DepartureTerminal     string
	ArrivalTerminal       string
	Equipment             string
	CabinClass            string
	BookingClass          string
	FareBasis             string
	FlightsMiles          int
	AwardFare             bool
	Duration              int
	LayoverDuration       int
	SubjectToGovtApproval bool
	SegmentRef            string
	Bookings              []Booking
}

type Booking struct {
	ID              uint `gorm:"primaryKey"`
	PassengerID     uint `gorm:"not null"`
	SeatID          uint `gorm:"not null;index:idx_seat_segment,unique"`
	FlightSegmentID uint `gorm:"not null;index:idx_seat_segment,unique"`
	BookedAt        time.Time
}
