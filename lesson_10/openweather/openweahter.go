package openweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RCoordinates struct {
	Lon float32
	Lat float32
}

type RWeather struct {
	Id          int32
	Main        string
	Description string
	Icon        string
}

type RMain struct {
	Temp       float32
	Feels_like float32
	Temp_min   float32
	Temp_max   float32
	Pressure   int32
	Humidity   int32
	Sea_level  int32
	Grnd_level int32
}

type RWind struct {
	Speed float32
	Deg   int32
	Gust  float32
}

type RClouds struct {
	All int32
}

type RSys struct {
	Type    int32
	Id      int32
	Country string
	Sunrise int32
	Sunset  int32
}

type Response struct {
	Coord      RCoordinates
	Weather    []RWeather
	Base       string
	Main       RMain
	Visibility int32
	Wind       RWind
	Clouds     RClouds
	Dt         int32
	Sys        RSys
	Timezone   int32
	Id         int32
	Name       string
	Cod        int32
}

func GetWeather(city string) (Response, error) {
	api_key := os.Getenv("OPENWEATHER_TOKEN")
	api_url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, api_key)

	client := &http.Client{}
	req, err := http.NewRequest("GET", api_url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	response := Response{}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		fmt.Printf("Cant parse weather response. \n Body: \n %s \n error %v", string(bodyBytes), err)
	}

	return response, err
}
