package cmd

import (
	"fmt"
	"log"
	"sort"

	"datadog.com/se-interview/api"
	"github.com/spf13/cobra"
)

// findCheapest represents the findCheapest command
var findCheapest = &cobra.Command{
	Use:   "findCheapest",
	Short: "Finds cheapest flight in a given window",
	Long: `
		Given a departure and arrival date, as well as a 3-character departure 
		and destination	city code (e.g.: JFK, LHR, DEN), the service will query 
		available flights, and find the cheapest flight for a given timeframe. 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		dep, err := cmd.Flags().GetString("dep")
		if dep == "" || err != nil {
			log.Fatalf("departure city (-dep, -d) is required")
		}

		dest, _ := cmd.Flags().GetString("dest")
		if dest == "" {
			log.Fatalf("destination city (-dest, -a) is required")
		}

		out, _ := cmd.Flags().GetString("out")
		if out == "" {
			log.Fatalf("outbound date (-out, -o) is required")
		}

		ret, _ := cmd.Flags().GetString("ret")
		if ret == "" {
			log.Fatalf("retrun date (-ret, -r) is required")
		}

		result := api.GetFlights(dep, dest, out, ret)

		sort.Slice(result.BestFlights, func(a, b int) bool {
			return result.BestFlights[a].Price < result.BestFlights[b].Price
		})

		fmt.Print("\n\nFLIGHT INFORMATION:\n\n")

		printFlightInfo(result)
	},
}

func init() {
	RootCmd.AddCommand(findCheapest)

	findCheapest.PersistentFlags().StringP("dep", "d", "", "A departure city for a flight")
	findCheapest.PersistentFlags().StringP("dest", "a", "", "A destination city for a flight")
	findCheapest.PersistentFlags().StringP("out", "o", "", "A departure date")
	findCheapest.PersistentFlags().StringP("ret", "r", "", "A return date")
}
