package cmd

import (
	"fmt"
	"time"

	"datadog.com/se-interview/model"
)

func printFlightInfo(result model.Result) {
	for _, flight := range result.BestFlights[0].Flights {

		fmt.Print(
			"\tDEPARTURE AIRPORT INFO:\n",
			"\t- Aiport Name:\t\t", flight.DepartureAirport.Name,
			"\n\t- Aiport ID:\t\t", flight.DepartureAirport.ID,
			"\n\t- Departure Time:\t", flight.DepartureAirport.Time,
			"\n\n\tARRIVAL AIRPORT INFO:\n",
			"\t- Aiport Name:\t\t", flight.ArrivalAirport.Name,
			"\n\t- Aiport ID:\t\t", flight.ArrivalAirport.ID,
			"\n\t- Departure Time:\t", flight.ArrivalAirport.Time,
			"\n\n\n")
	}

	fmt.Println(
		"\n\nTOTAL PRICE:\t\t\t\t", result.BestFlights[0].Price,
		"\nFLIGHT DURATION:\t\t\t", time.Duration(float64(result.BestFlights[0].TotalDuration)*float64(time.Minute)),
	)
}
