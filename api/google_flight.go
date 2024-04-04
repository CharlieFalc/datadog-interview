package api

import (
	"encoding/json"

	"datadog.com/se-interview/model"
	g "github.com/serpapi/google-search-results-golang"
)

func GetFlights(dep, dest, out, ret string) model.Result {
	var result model.Result

	// REAL DATA START
	parameter := map[string]string{
		"engine":        "google_flights",
		"departure_id":  dep,
		"arrival_id":    dest,
		"outbound_date": out,
		"return_date":   ret,
		"currency":      "USD",
		"hl":            "en",
		"api_key":       "730df22274fe208eb50f8b5a841f957a9fdfdb07dd382920831aab4d978382bc",
	}
	search := g.NewGoogleSearch(parameter, "730df22274fe208eb50f8b5a841f957a9fdfdb07dd382920831aab4d978382bc")
	searchJSON, _ := search.GetJSON()
	jsonStr, _ := json.Marshal(searchJSON)
	json.Unmarshal(jsonStr, &result)
	// REAL DATA END

	// MOCK DATA START
	// jsonFile, err := os.Open("mock_json.json")
	// if err != nil {
	// 	fmt.Println("failed to open mock_json.json")
	// }
	// byteVal, _ := io.ReadAll(jsonFile) // mock data
	// json.Unmarshal(byteVal, &result)
	// MOCK DATA END

	return result
}
