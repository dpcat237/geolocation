package service

import (
	"gitlab.com/dpcat237/geolocation/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	locationTestCases = []locationTestCase{
		{
			Location: model.LocationRow{
				IPAddress:    "1.1.1.1",
				CountryCode:  "GB",
				Country:      "United Kingdom",
				City:         "London",
				LatitudeStr:  "-44.02628376337595",
				LongitudeStr: "38.69732042689665",
			},
			Valid: true,
		},
		{
			Location: model.LocationRow{
				IPAddress:    "123.112.144.168",
				CountryCode:  "LO",
				Country:      "Lornt",
				City:         "Todfjg",
				LatitudeStr:  "-44.045628376337595",
				LongitudeStr: "38.697089042689665",
			},
			Valid: true,
		},
		{
			Location: model.LocationRow{
				IPAddress:    "13.11.14.168",
				CountryCode:  "LO",
				Country:      "Lornt",
				City:         "Todfjg",
				LatitudeStr:  "0",
				LongitudeStr: "38.697089042689665",
			},
			Valid: true,
		},
		{
			Location: model.LocationRow{
				IPAddress:    "131.14.168",
				CountryCode:  "LO",
				Country:      "Lornt",
				City:         "Todfjg",
				LatitudeStr:  "0",
				LongitudeStr: "38.697089042689665",
			},
			Valid: false,
		},
		{
			Location: model.LocationRow{
				IPAddress:    "13.1154.14.168",
				CountryCode:  "LO",
				Country:      "Lornt",
				City:         "Todfjg",
				LatitudeStr:  "0",
				LongitudeStr: "38.697089042689665",
			},
			Valid: false,
		},
		{
			Location: model.LocationRow{
				IPAddress:    "346456546",
				CountryCode:  "LO",
				Country:      "Lornt",
				City:         "Todfjg",
				LatitudeStr:  "0",
				LongitudeStr: "38.697089042689665",
			},
			Valid: false,
		},
		{
			Location: model.LocationRow{
				IPAddress:    "",
				CountryCode:  "LO",
				Country:      "Lornt",
				City:         "Todfjg",
				LatitudeStr:  "0",
				LongitudeStr: "38.697089042689665",
			},
			Valid: false,
		},
		{
			Location: model.LocationRow{
				IPAddress:    "13.11.14.168",
				CountryCode:  "",
				Country:      "Lornt",
				City:         "Todfjg",
				LatitudeStr:  "0",
				LongitudeStr: "38.697089042689665",
			},
			Valid: false,
		},
		{
			Location: model.LocationRow{
				IPAddress:    "13.11.14.168",
				CountryCode:  "LO",
				Country:      "Lornt",
				City:         "Todfjg",
				LatitudeStr:  "0",
				LongitudeStr: "38.697089042689t65",
			},
			Valid: false,
		},
		{
			Location: model.LocationRow{
				IPAddress:    "13.11.14.168",
				CountryCode:  "LO",
				Country:      "",
				City:         "Todfjg",
				LatitudeStr:  "0",
				LongitudeStr: "38.697089042689665",
			},
			Valid: false,
		},
	}
)

type countryServiceTest struct{}
type locationRepositoryTest struct{}

func (srv *countryServiceTest) CreateCountry(iso, nm string) (model.Country, model.Error) {
	return model.Country{}, model.Error{}
}
func (srv *countryServiceTest) GetCountries() (model.Countries, model.Error) {
	return nil, model.Error{}
}

func (rps *locationRepositoryTest) GetLastLocationByIP(ip string) (model.Location, error) {
	return model.Location{}, nil
}
func (rps *locationRepositoryTest) CreateLocationsLocationsBulk(los []*model.Location) error {
	return nil
}

type locationTestCase struct {
	Location model.LocationRow
	Valid    bool
}

func TestEventDTO_CalculatePartnerProfit(t *testing.T) {
	var couSrv *countryServiceTest
	var locRps *locationRepositoryTest
	locSrv := newLocation(couSrv, locRps)
	for _, tc := range locationTestCases {
		locs := locSrv.sanitizeLocationsData([]model.LocationRow{tc.Location})
		valid := len(locs) == 1
		assert.Equal(t, tc.Valid, valid)
	}
}
