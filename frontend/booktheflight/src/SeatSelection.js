import { useEffect, useState } from 'react';
import './SeatSelection.css';

const SeatSelection = () => {
  const [selectedSeats, setSelectedSeats] = useState([]);
  const [data, setSeatData] = useState(null);

  useEffect(() => {
    const params = new URLSearchParams({
      passenger_id: process.env.REACT_APP_PASSENGER_UUID_TEST,
      flight_id: process.env.REACT_APP_FLIGHT_UUID_TEST,
    });

    fetch(`${process.env.REACT_APP_API_BASE_URL}/v1/seat?${params.toString()}`)
      .then((res) => {
        if (!res.ok) {
          throw new Error('Fetch failed: ' + res.status);
        }
        return res.json();
      })
      .then((data) => {
        setSeatData(data);
      })
      .catch((err) => console.error('Fetch error:', err));
  }, []);

  if (!data) return <p>Loading...</p>;

  // Only allow one selected seat at a time
  const toggleSeat = (seatCode) => {
    setSelectedSeats((prevSelected) =>
      prevSelected.includes(seatCode) ? [] : [seatCode]
    );
  };

  const allSeats =
    data.seatsItineraryParts?.[0]
      ?.segmentSeatMaps?.[0]
      ?.passengerSeatMaps?.[0]
      ?.seatMap?.cabins
      ?.flatMap((cabin) => cabin.seatRows)
      ?.flatMap((row) => row.seats) || [];
      
    const segment = 
    data.seatsItineraryParts?.[0]
      ?.segmentSeatMaps?.[0].segment

    const passenger = 
     data.seatsItineraryParts?.[0]
      ?.segmentSeatMaps?.[0]
      ?.passengerSeatMaps?.[0].passenger

  return (
    <div className="seat-selection-container">
      {/* Flight + Passenger Info */}
      <div className="info-row">
        <div className="flight-details">
          <h2>Flight Information</h2>
          <p><strong>Flight Number:</strong> {segment.flight.airlineCode}{segment.flight.flightNumber} </p>
          <p><strong>Route:</strong> {segment.origin} ({segment.flight.departureTerminal}) - {segment.destination} ({segment.flight.arrivalTerminal})</p>
          <p><strong>Departure:</strong> {segment.departure} </p>
          <p><strong>Arrival:</strong> {segment.arrival} </p>
        </div>

        <div className="passenger-details">
          <h2>Passenger Information</h2>
          <p><strong>Name:</strong> {passenger.passengerDetails.lastName}, {passenger.passengerDetails.firstName}</p>
          <p><strong>Date of Birth:</strong> {passenger.passengerInfo.dateOfBirth} </p>
          <h3><strong>Frequent Flyer</strong> </h3>
          {passenger.preferences.frequentFlyer.map((ff, fi)=>{
            return (
                <div key={fi}> 
                    <p> {ff.airline} - {ff.number} - Tier {ff.tierNumber} </p>
                </div>
            );
          })}
        </div>
      </div>

    {/* Selected Seat Info */}
      {selectedSeats.length > 0 && (
        <div className="seat-info">
          {selectedSeats.map((seatCode, si) => {
            const seatDetails = allSeats.find((seat) => seat.code === seatCode);
            return (
              <div className="seat-details" key={si}>
                <h2>Selected Seat: {seatDetails?.code}</h2>
                <div className="seat-price">
                  <h3>Price</h3>
                    {seatDetails.prices.alternatives.map((p, pi)=>{
                    return (
                        <div key={pi}>
                            <p> <strong> {p[0].amount} {p[0].currency} </strong> </p>
                        </div>
                    )
                  })}
                </div>
              </div>
            );
          })}
        </div>
      )}

      {/* Seat Map */}
      <div className="seat-map-container">
        <center><h2>Seat Map</h2></center>
        {data.seatsItineraryParts[0].segmentSeatMaps[0].passengerSeatMaps[0].seatMap.cabins.map((cabin, ci) => (
          <div key={ci}>
            <h3 className="text-lg font-medium mb-4"> <center>
              {cabin.deck} Deck - Rows {cabin.firstRow} to {cabin.lastRow} </center>
            </h3>
            {cabin.seatRows.map((row, ri) => (
              <div className="seat-row" key={ri}>
                {row.seats.map((seat, si) => {
                  const code = seat.storefrontSlotCode;

                  if (code === "BLANK") {
                    return (
                      <div
                        key={si}
                        className={`seat blank-seat ${seat.slotCharacteristics?.includes("LEFT_SIDE") ? 'left' : ''} ${seat.slotCharacteristics?.includes("RIGHT_SIDE") ? 'right' : ''}`}
                      />
                    );
                  }

                  if (code === "AISLE") {
                    return <div key={si} className="seat aisle-seat" />;
                  }

                  if (code === "BULKHEAD") {
                    return <div key={si} className="seat bulkhead-seat" />;
                  }

                  if (code === "WING") {
                    return <div key={si} className="seat wing-area" />;
                  }

                  return (
                    <div
                      key={si}
                      className={`seat ${seat.available ? 'available' : 'unavailable'} ${
                        selectedSeats.includes(seat.code) ? 'selected' : ''
                      }`}
                      onClick={() => seat.available && toggleSeat(seat.code)}
                    >
                      {seat.code}
                    </div>
                  );
                })}
              </div>
            ))}
          </div>
        ))}
      </div>
    </div>
  );
};

export default SeatSelection;
