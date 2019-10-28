package controller

import (
	"gitlab.com/dpcat237/geolocation/logger"
	"gitlab.com/dpcat237/geolocation/service"
)

// Collector defines controllers
type Collector struct {
	LocCtr LocationController
}

// InitCollector initializes collector of controllers for gRPC
func InitCollector(logg logger.Logger, sCll service.Collector) Collector {
	return Collector{
		LocCtr: newLocation(logg, sCll.LocSrv),
	}
}
