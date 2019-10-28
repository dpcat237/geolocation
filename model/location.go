package model

import (
	"time"

	"gitlab.com/dpcat237/geomicroservices/geolocation"
)

// LocationTable defines name of locations table on database
const LocationTable = "locations"

// Location defines location struct
type Location struct {
	ID        uint32 `gorm:"primary_key"`
	IPAddress string `gorm:"column:ip_address"`
	Latitude  float64
	Longitude float64
	CountryID uint32
	Country   Country `gorm:"association_autoupdate:false;association_autocreate:false"`
	City      string
	CreatedAt int64
}

// TableName return table name of Location struct for gorm
func (Location) TableName() string {
	return LocationTable
}

// DefineLocation creates Location defining past parameters
func DefineLocation(ipAd string, lat, lon float64, couID uint32, city string) Location {
	return Location{
		IPAddress: ipAd,
		Latitude:  lat,
		Longitude: lon,
		City:      city,
		CountryID: couID,
		CreatedAt: time.Now().UTC().Unix(),
	}
}

// ToGRPC converts Location struct to gRPC standardized struct
func (loc Location) ToGRPC() *geolocation.Location {
	return &geolocation.Location{
		Id:        loc.ID,
		IpAddress: loc.IPAddress,
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
		Country:   loc.Country.ToGRPC(),
		City:      loc.City,
	}
}
