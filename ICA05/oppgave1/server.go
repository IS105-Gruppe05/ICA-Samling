package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path"
)
//This struct was generated from the BF4 API available at http://api.bf4stats.com/api/onlinePlayers?output=json
//We used the JSON to GO tool available at https://mholt.github.io/json-to-go/
type GeneratedGaming struct {
	Pc struct {
		Label  string `json:"label"`
		Count  int    `json:"count"`
		Peak24 int    `json:"peak24"`
	} `json:"pc"`
	Ps3 struct {
		Label  string `json:"label"`
		Count  int    `json:"count"`
		Peak24 int    `json:"peak24"`
	} `json:"ps3"`
	Xbox struct {
		Label  string `json:"label"`
		Count  int    `json:"count"`
		Peak24 int    `json:"peak24"`
	} `json:"xbox"`
	Xone struct {
		Label  string `json:"label"`
		Count  int    `json:"count"`
		Peak24 int    `json:"peak24"`
	} `json:"xone"`
	Ps4 struct {
		Label  string `json:"label"`
		Count  int    `json:"count"`
		Peak24 int    `json:"peak24"`
	} `json:"ps4"`
}
//This struct was generated from the Apuxi Weather API available at http://api.apixu.com/v1/current.json?key=3cc8e1ca15b24b72a60121151172803&q=London
//We used the JSON to GO tool available at https://mholt.github.io/json-to-go/
type GeneratedWeather struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
	} `json:"current"`
}
//This struct was generated from a sample of the OpenWeatherMap API available at http://samples.openweathermap.org/data/2.5/weather?zip=94040%2Cus&appid=b1b15e88fa797225412429c1c50c122a1
//We used the JSON to GO tool available at https://mholt.github.io/json-to-go/
type GeneratedWeather2 struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

//Here we initialize the server at port 8001
func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8001", nil)

}
//The workhorse. Foo reads, decodes and writes JSON values to desired HTML templates.
func foo(w http.ResponseWriter, r *http.Request) {
	//Gets URL values and checks for a response
	url := "http://api.apixu.com/v1/current.json?key=3cc8e1ca15b24b72a60121151172803&q=London"
	url2 := "http://api.bf4stats.com/api/onlinePlayers?output=json"
	url3 := "http://samples.openweathermap.org/data/2.5/weather?zip=94040%2Cus&appid=b1b15e88fa797225412429c1c50c122a1"
	resp, err := http.Get(url)
	resp2, err := http.Get(url2)
	resp3, err := http.Get(url3)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	//Pass JSON to memory and decode according to desired struct
	defer resp.Body.Close()
	var w1 GeneratedWeather
	err = json.NewDecoder(resp.Body).Decode(&w1)
	if err != nil {
		log.Fatal(err)
	}
	//Same as above, but for the GeneratedGaming struct
	defer resp2.Body.Close()
	var w2 GeneratedGaming
	err = json.NewDecoder(resp2.Body).Decode(&w2)
	if err != nil {
		log.Fatal(err)
	}
	//Decode GeneratedWeather2 struct
	defer resp3.Body.Close()
	var w3 GeneratedWeather2
	err = json.NewDecoder(resp3.Body).Decode(&w3)
	if err != nil {
		log.Fatal(err)
	}
	//Write JSON values to desired HTML template
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, w1); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//Write JSON values to desired HTML template
	fp3 := path.Join("templates", "index3.html")
	tmpl3, err := template.ParseFiles(fp3)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl3.Execute(w, w3); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//Write JSON values to desired HTML template
	fp2 := path.Join("templates", "index2.html")
	tmpl2, err := template.ParseFiles(fp2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl2.Execute(w, w2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}


}
