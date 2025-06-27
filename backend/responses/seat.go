package responses

type SeatsItineraryParts struct {
	SegmentSeatMaps `json:"segmentSeatMaps"`
}

type SegmentSeatMaps struct {
	PassengerSeatMaps []PassengerSeatMap `json:"segmentSeatMaps"`
	Segment           `json:"segment"`
}

type SeatMap struct {
	RowsDisabledCauses []any   `json:"rowsDisabledCauses"`
	Aircraft           string  `json:"aircraft"`
	Cabins             []Cabin `json:"cabins"`
}

type Seat struct {
	StorefrontSlotCode  string   `json:"storefrontSlotCode"`
	Available           bool     `json:"available"`
	Code                string   `json:"code"`
	Designations        []any    `json:"designations"`
	Entitled            bool     `json:"entitled"`
	FeeWaived           bool     `json:"feeWaived"`
	EntitledRuleID      string   `json:"entitledRuleId"`
	FeeWaivedRuleID     string   `json:"feeWaivedRuleId"`
	SeatCharacteristics []string `json:"seatCharacteristics"`
	Limitations         []any    `json:"limitations"`
	RefundIndicator     string   `json:"refundIndicator"`
	FreeOfCharge        bool     `json:"freeOfCharge"`
	Prices              struct {
		Alternatives []Alternative `json:"alternatives"`
	} `json:"prices"`
	Taxes struct {
		Alternatives []Alternative `json:"alternatives"`
	} `json:"taxes"`
	Total struct {
		Alternatives []Alternative `json:"alternatives"`
	} `json:"total"`
	OriginallySelected     bool     `json:"originallySelected"`
	RawSeatCharacteristics []string `json:"rawSeatCharacteristics"`
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
	Preference          `json:"preferences"`
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
	Phones      []any    `json:"phones"`
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
	SpecialPreference `json:"specialPreferences"`
	FrequentFlyer     []FrequentFlyer `json:"frequentFlyer"`
}

type FrequentFlyer struct {
	Airline    string `json:"airline"`
	Number     string `json:"number"`
	TierNumber int    `json:"tierNumber"`
}

type SpecialPreference struct {
	MealPreference               string `json:"mealPreference"`
	SeatPreference               string `json:"seatPreference"`
	SpecialRequests              []any  `json:"specialRequests"`
	SpecialServiceRequestRemarks []any  `json:"specialServiceRequestRemarks"`
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
	FlightNumber          int    `json:"flightNumber"`
	OperatingFlightNumber int    `json:"operatingFlightNumber"`
	AirlineCode           string `json:"airlineCode"`
	OperatingAirlineCode  string `json:"operatingAirlineCode"`
	StopAirports          []any  `json:"stopAirports"`
	DepartureTerminal     string `json:"departureTerminal"`
	ArrivalTerminal       string `json:"arrivalTerminal"`
}

type GetSeat struct {
	SeatsItineraryParts `json:"seatsItineraryParts"`
	SelectedSeats       []any `json:"selectedSeats"`
}
