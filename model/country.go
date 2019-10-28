package model

import "gitlab.com/dpcat237/geomicroservices/geolocation"

// CountryTable defines name of countries table on database
const CountryTable = "countries"

// Country defines country struct
type Country struct {
	ID   uint32 `gorm:"primary_key"`
	ISO  string
	Name string
}

// Countries defines slice of Country struct
type Countries []Country

// TableName return table name of Country struct for gorm
func (Country) TableName() string {
	return CountryTable
}

// GetCountriesISO return slice of countries ISO code
func (cs Countries) GetCountriesISO() []string {
	var isos []string
	for _, c := range cs {
		isos = append(isos, c.ISO)
	}
	return isos
}

// GetMappedByISO converts slice of countries to map of countries ordered by ISO and returns it
func (cos Countries) GetMappedByISO() map[string]Country {
	cosMap := make(map[string]Country)
	for _, co := range cos {
		cosMap[co.ISO] = co
	}
	return cosMap
}

// GetCountryFromMapByISO extracts country from countries map by ISO code
func (co Country) GetCountryFromMapByISO(cosMap map[string]Country, iso string) Country {
	if coDt, exists := cosMap[iso]; exists {
		return coDt
	}
	return Country{}
}

// ToGRPC converts Country struct to gRPC standardized struct
func (co Country) ToGRPC() *geolocation.Country {
	return &geolocation.Country{
		Id:   co.ID,
		Iso:  co.ISO,
		Name: co.Name,
	}
}
