package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IP struct {
	IP string `json:"ip"`
}

type LATLON struct {
	LATITUDE  float64 `json:"lat"`
	LONGITUDE float64 `json:"lon"`
	City      string  `json:"city"`
}

type WEATHERDATA struct {
	Current struct {
		Temperature   float64 `json:"temperature_2m"`
		WindSpeed     float64 `json:"wind_speed_10m"`
		WindGustSpeed float64 `json:"wind_gusts_10m"`
		Precipitation float64 `json:"precipitation"`
	} `json:"current"`
}

func main() {
	fmt.Println("IP LOCATION WEATHER TOOL")

	var ipthing IP
	ipreq, _ := http.Get("https://api.ipify.org?format=json")
	defer ipreq.Body.Close()

	ipaftrio, _ := ioutil.ReadAll(ipreq.Body)
	// fmt.Println(string(ipaftrio)) // for debugging, you can remove this

	err := json.Unmarshal(ipaftrio, &ipthing)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("IP: ", ipthing.IP)
	requestthing := string("http://ip-api.com/json/" + ipthing.IP)
	iplocreq, _ := http.Get(requestthing)
	defer iplocreq.Body.Close()
	stringiplocreq, _ := ioutil.ReadAll(iplocreq.Body)
	//fmt.Println(string(stringiplocreq))
	var latlonthing LATLON

	err = json.Unmarshal(stringiplocreq, &latlonthing)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(latlonthing.LATITUDE, latlonthing.LONGITUDE)
	// we now have the users latitude and longitude based on the user ip, stored as latlonthing.LATITUDE and latlonthing.LONGITUDE as float64

	finishedweatherrequest := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%.6f&longitude=%.6f&current=temperature_2m,wind_speed_10m,wind_gusts_10m,precipitation&timezone=auto",
		latlonthing.LATITUDE,
		latlonthing.LONGITUDE,
	)
	wthrreq, _ := http.Get(finishedweatherrequest)
	stringwthrreq, _ := ioutil.ReadAll(wthrreq.Body)
	defer wthrreq.Body.Close()
	//fmt.Println(string(stringwthrreq))
	var weatherjsontable WEATHERDATA
	err = json.Unmarshal(stringwthrreq, &weatherjsontable)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The weather in ", latlonthing.City, " is:")
	fmt.Println("Temperature: ", weatherjsontable.Current.Temperature, " Celsius")
	fmt.Println("Wind speed: ", weatherjsontable.Current.WindSpeed, " KM/H")
	fmt.Println("Wind gust speed: ", weatherjsontable.Current.WindGustSpeed, " KM/H")
	fmt.Println("Precipitation (Rain/Snow amount): ", weatherjsontable.Current.Precipitation, " mm")
	fmt.Println("LATITUDE: ", latlonthing.LATITUDE)
	fmt.Println("LONGITUDE:", latlonthing.LONGITUDE)
}
