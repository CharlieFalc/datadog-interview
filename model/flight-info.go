package model

type Result struct {
	BestFlights []struct {
		Flights       []Flight `json:"flights"`
		TotalDuration int      `json:"total_duration"`
		Price         int      `json:"price"`
		Type          string   `json:"type"`
		Extensions    []string `json:"extensions"`
	} `json:"best_flights"`
	OtherFlights []struct {
		Flights  []Flight `json:"flights"`
		Layovers []struct {
			Duration int    `json:"duration"`
			Name     string `json:"name"`
			ID       string `json:"id"`
		} `json:"layovers,omitempty"`
		TotalDuration  int      `json:"total_duration"`
		Price          int      `json:"price"`
		Type           string   `json:"type"`
		AirlineLogo    string   `json:"airline_logo"`
		Extensions     []string `json:"extensions"`
		DepartureToken string   `json:"departure_token"`
	} `json:"other_flights"`
	PriceInsights struct {
		LowestPrice       int     `json:"lowest_price"`
		PriceLevel        string  `json:"price_level"`
		TypicalPriceRange []int   `json:"typical_price_range"`
		PriceHistory      [][]int `json:"price_history"`
	} `json:"price_insights"`
}

type Flight struct {
	DepartureAirport AirportInfo `json:"departure_airport"`
	ArrivalAirport   AirportInfo `json:"arrival_airport"`
	Duration         int         `json:"duration"`
	Airplane         string      `json:"airplane"`
	Airline          string      `json:"airline"`
	TravelClass      string      `json:"travel_class"`
	FlightNumber     string      `json:"flight_number"`
	Extensions       []string    `json:"extensions"`
	Overnight        bool        `json:"overnight"`
}

type AirportInfo struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Time string `json:"time"`
}
