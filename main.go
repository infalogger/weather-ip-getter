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
}

func main() {
	fmt.Println("hello")

	var ipthing IP
	ipreq, _ := http.Get("https://api.ipify.org?format=json")
	defer ipreq.Body.Close()

	ipaftrio, _ := ioutil.ReadAll(ipreq.Body)
	fmt.Println(string(ipaftrio)) // for debugging, you can remove this

	err := json.Unmarshal(ipaftrio, &ipthing)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ipthing.IP)
	requestthing := string("http://ip-api.com/json/" + ipthing.IP)
	iplocreq, _ := http.Get(requestthing)
	defer iplocreq.Body.Close()
	stringiplocreq, _ := ioutil.ReadAll(iplocreq.Body)
	fmt.Println(string(stringiplocreq))
	var latlonthing LATLON

	err = json.Unmarshal(stringiplocreq, &latlonthing)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(latlonthing.LATITUDE, latlonthing.LONGITUDE)
	// we now have the users latitude and longitude based on the user ip, stored as latlonthing.LATITUDE and latlonthing.LONGITUDE as float64

}
