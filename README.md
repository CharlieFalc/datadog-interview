# Datadog Interview

## Build/Deployment
To build the following code, simply execute:
```go build .```

Then for execution, leverage the built package name `se-interview -h` to see a list of the available commands with the CLI tool. 

## General Usage
As defined in the se-interview helper function, there are two available commands within the tool `findCheapest` and `findFastest`. Each command takes in a set of required flags to search for a flight:
```
  -d, --dep string    A departure city for a flight
  -a, --dest string   A destination city for a flight
  -h, --help          help for findFastest
  -o, --out string    A departure date
  -r, --ret string    A return date
```

Execute the `findFastest` or `findCheapest` flight with the above flags to execute a query for the cheapest flight. The tool mainly leverages SerpApi to query Google for flight data, refer to https://serpapi.com/google-flights-results for documentation on Serp
