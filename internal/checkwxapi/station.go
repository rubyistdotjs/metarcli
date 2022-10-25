package checkwxapi

import (
	"fmt"
	"os"
)

type Station struct {
	Type     string
	Name     string
	Iata     string
	IcaoCode string `json:"icao"`
	City     string
	State    struct {
		Code string
		Name string
	}
	Country struct {
		Code string
		Name string
	}
	Elevation struct {
		Feet   int16
		Meters int16
	}
	Latitude  StationGeoPoint
	Longitude StationGeoPoint
	Geometry  struct {
		Coordinates []float64
		Type        string
	}
	Timezone struct {
		Dst  int
		Gmt  string
		Tzid string
		Zone string
	}
}

type StationGeoPoint struct {
	Decimal float64
	Degrees string
}

type getStationRes struct {
	Data    []Station
	Results int
}

func (c *Client) RetrieveStations(icaoCodes []string) map[string]Station {
	var body getStationRes
	err := c.get("station/", icaoCodes, &body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result := make(map[string]Station, len(icaoCodes))

	for _, station := range body.Data {
		result[station.IcaoCode] = station
	}

	return result
}
