package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//func ReverseGeocode(lat, lon float64) (string, error) {
//geocoder := new(geo.GoogleGeocoder)
//coords := geo.NewPoint(lat, lon)
//return geocoder.ReverseGeocode(coords)
//}

type GeocodeResult struct {
	Name string `json:"name"`
}

type GooglePlacesResponse struct {
	Status  string          `json:"status"`
	Results []GeocodeResult `json:"results"`
}

func ReverseGeocode(lat, lon float64) (string, error) {
	apiKey := "AIzaSyBvqVz9o3MZUQLeBgZgQ-ATAGSArovTV9o"
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%f,%f&radius=1000&key=%s", lat, lon, apiKey)
	req, err := http.Get(url)
	decoder := json.NewDecoder(req.Body)

	var response GooglePlacesResponse
	if err = decoder.Decode(&response); err != nil {
		return "", err
	}

	if response.Status != "OK" || len(response.Results) == 0 {
		return "", errors.New("Geocode request failed")
	}

	return response.Results[0].Name, nil
}
