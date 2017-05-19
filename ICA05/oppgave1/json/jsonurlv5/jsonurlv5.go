package jsonurlv5

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"log"
	"strings"
)

func JsonUrlV5() {

	url := "http://api.apixu.com/v1/current.json?key=5aac958781c64de0b43121244172903&q=Oslo"
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

	type Location struct {
		Name string
		Lat float64
		Lon float64
	}
	type Measurements struct {
		Temp_c   float64
		Humidity float64
		Wind_kph float64
	}
	type Weather struct {
		Location Location
		Current  Measurements
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var w Weather
		if err := dec.Decode(&w); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Location: %s\n", w.Location.Name)
		fmt.Printf("Coordinates are: longitude %.2f and latitude %.2f\n", w.Location.Lon, w.Location.Lat)
		fmt.Printf("Temperature in Celcius: %.2f\n", w.Current.Temp_c)
		fmt.Printf("Humidity: %.f\n", w.Current.Humidity)
		fmt.Printf("Wind kph: %.2f\n\n", w.Current.Wind_kph)
	}
}
