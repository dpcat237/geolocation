package controller

import (
	"net/http"
	"time"

	"gitlab.com/dpcat237/geomicroservices/geolocation"
	"golang.org/x/net/context"

	"gitlab.com/dpcat237/geolocation/logger"
	"gitlab.com/dpcat237/geolocation/service"
)

// LocationController defines required methods for Location controller
type LocationController interface {
	GetLocationByID(ctx context.Context, req *geolocation.GetLocationByIDRequest) (*geolocation.GetLocationByIDResponse, error)
}

// locationController defines required services for Location controller
type locationController struct {
	logg   logger.Logger
	locSrv service.LocationService
}

// newLocation initializes Location controller for gRPC
func newLocation(logg logger.Logger, locSrv service.LocationService) *locationController {
	return &locationController{logg: logg, locSrv: locSrv}
}

// GetLocationByID returns location details for requested IP
func (ctr *locationController) GetLocationByID(ctx context.Context, req *geolocation.GetLocationByIDRequest) (*geolocation.GetLocationByIDResponse, error) {
	status := uint32(http.StatusOK)
	act := "location.get_by_ip"
	var errMsg string
	defer ctr.logg.RequestEnd(time.Now(), act, &status, &errMsg)

	loc, er := ctr.locSrv.GetLocationByIP(req.IpAddress)
	if er.IsError() {
		status = er.Status
		errMsg = er.GetLogMessage()
		return &geolocation.GetLocationByIDResponse{
			Code:    uint32(status),
			Message: er.Message,
		}, nil
	}
	return &geolocation.GetLocationByIDResponse{
		Code:     uint32(status),
		Location: loc.ToGRPC(),
	}, nil
}
