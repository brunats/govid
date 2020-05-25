package nowsh

import "time"

type countryInfo struct {
	Country   string    `json:"country"`
	Cases     int       `json:"cases"`
	Confirmed int       `json:"confirmed"`
	Recovered int       `json:"recovered"`
	Deaths    int       `json:"deaths"`
	UpdatedAt time.Time `json:"updated_at"`
}
