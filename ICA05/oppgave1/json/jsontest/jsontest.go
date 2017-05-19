package jsontest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func ExDecodeMine() {
	const jsonStream = `{
    "coord":{"lon":-122.08,"lat":37.39},
"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10n"}],
"base":"stations",
"main":{"temp":277.14,"pressure":1025,"humidity":86,"temp_min":275.15,"temp_max":279.15},
"visibility":16093,
"wind":{"speed":1.67,"deg":53.0005},
"clouds":{"all":1},
"dt":1485788160,
"sys":{"type":1,"id":471,"message":0.0116,"country":"US","sunrise":1485789140,"sunset":1485826300},
"id":5375480,
"name":"Mountain View",
"cod":200
    }`
	type Coordinates struct {
		Lon float64
		Lat float64
	}
	type Measurements struct {
		Temp     float64
		Pressure float64
		Humidity float64
		Temp_min float64
		Temp_max float64
		Wind     float64
		Speed    float64
	}
	type Weather struct {
		Coord Coordinates
		Main  Measurements
		Wind  Measurements
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var w Weather
		if err := dec.Decode(&w); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Coordinates are: longitude %.2f and latitude %.2f\n", w.Coord.Lon, w.Coord.Lat)
		fmt.Printf("Temperature: %.f\n", w.Main.Temp)
		fmt.Printf("Test: %.f\n", w.Main.Humidity)
		fmt.Printf("Wind speed: %.2f\n", w.Wind.Speed)

	}
}
