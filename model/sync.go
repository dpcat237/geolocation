package model

import "strings"

// LocationRow defines parameters to load from location CSV
type LocationRow struct {
	IPAddress    string `csv:"ip_address"`
	CountryCode  string `csv:"country_code"`
	Country      string `csv:"country"`
	City         string `csv:"city"`
	LatitudeStr  string `csv:"latitude"`
	LongitudeStr string `csv:"longitude"`
	Latitude     float64
	Longitude    float64
}

// LocationRows defines slice of LocationRow struct
type LocationRows []LocationRow

// IsAllRequiredData checks if struct contains required data
func (lo LocationRow) IsAllRequiredData() bool {
	return lo.IPAddress != "" && lo.CountryCode != "" && lo.Country != ""
}

// GetCitySlug generates city slug from city name
func (lo LocationRow) GetCitySlug() string {
	return strings.Replace(strings.ToLower(lo.City), " ", "", -1)
}
