package responses

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

type SeatDetail struct {
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

type Seat struct {
	SlotCharacteristics []string `json:"slotCharacteristics,omitempty"`
	StorefrontSlotCode  string   `json:"storefrontSlotCode"`
	Available           bool     `json:"available"`
	Entitled            bool     `json:"entitled"`
	FeeWaived           bool     `json:"feeWaived"`
	FreeOfCharge        bool     `json:"freeOfCharge"`
	OriginallySelected  bool     `json:"originallySelected"`
	*SeatDetail
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
