package service

import (
	"gitlab.com/dpcat237/geolocation/repository/sql"
)

// Collector defines services
type Collector struct {
	LocSrv LocationService
}

// Init initializes services and required repositories
func Init(dbCl sql.DatabaseCollector) Collector {
	// Initialize repositories
	couRps := sql.NewCountry(dbCl.GetDatabase())
	locRps := sql.NewLocation(dbCl.GetDatabase())

	// Initialize services
	couSrv := newCountry(couRps)
	locSrv := newLocation(couSrv, locRps)

	return Collector{
		LocSrv: locSrv,
	}
}
