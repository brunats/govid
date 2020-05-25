package nowsh

type responseFormAll struct {
	CountryInfos []countryInfo `json:"data"`
}

type responseFormOne struct {
	CountryInfo countryInfo `json:"data"`
}
