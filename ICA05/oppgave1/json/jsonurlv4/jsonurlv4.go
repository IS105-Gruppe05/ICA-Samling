package jsonurlv4

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"log"
	"strings"
)

func JsonUrlV4() {

	url := "http://samples.openweathermap.org/data/2.5/weather?id=2172797&appid=b1b15e88fa797225412429c1c50c122a1"
	fmt.Printf("HTML code of %s \n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var jsonStream = string(html[:])

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
		fmt.Printf("Temperature in Kelvin: %.f\n", w.Main.Temp)
		fmt.Printf("Humidity: %.f\n", w.Main.Humidity)
		fmt.Printf("Wind speed: %.2f\n\n", w.Wind.Speed)

	}
}
