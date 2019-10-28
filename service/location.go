package service

import (
	"gitlab.com/dpcat237/geolocation/model"
	"gitlab.com/dpcat237/geolocation/repository/file"
	"gitlab.com/dpcat237/geolocation/repository/sql"
	"net"
	"strconv"
)

const (
	// locationsLocationsFile defines name of CSV file with locations data
	locationsLocationsFile = "geolocations.csv"
	// saveBunch define limit of saved location in a bunch
	saveBunch = 1000
)

// LocationService defines required methods for Location service
type LocationService interface {
	GetLocationByIP(ip string) (model.Location, model.Error)
	LoadLocations() (int, int, model.Error)
}

// locationService defines required services for Location service
type locationService struct {
	couSrv CountryService
	locRps sql.LocationRepository
}

// newLocation initializes Location service
func newLocation(couSrv CountryService, locRps sql.LocationRepository) *locationService {
	return &locationService{couSrv: couSrv, locRps: locRps}
}

// GetLocationByIP returns last locations details by IP address
func (srv locationService) GetLocationByIP(ip string) (model.Location, model.Error) {
	loc, err := srv.locRps.GetLastLocationByIP(ip)
	if err != nil && !sql.CheckNotFound(err) {
		return loc, model.NewErrorServer("Error getting location by IP").WithError(err)
	}
	return loc, model.NewErrorNil()
}

// LoadLocations loads locations details from CSV to database
func (srv locationService) LoadLocations() (int, int, model.Error) {
	var fltCount int
	var addCount int
	flCli, er := file.NewClient(locationsLocationsFile)
	if er.IsError() {
		return fltCount, addCount, er
	}
	defer flCli.GracefulClose()

	locsAll, er := flCli.GetLocations()
	if er.IsError() {
		return fltCount, addCount, er
	}
	if len(locsAll) == 0 {
		return fltCount, addCount, model.NewErrorNil()
	}
	locsRow := srv.sanitizeLocationsData(locsAll)
	if len(locsRow) == 0 {
		return fltCount, addCount, model.NewErrorNil()
	}
	fltCount = len(locsAll) - len(locsRow)

	cous, er := srv.couSrv.GetCountries()
	if er.IsError() {
		return fltCount, addCount, er
	}
	cousMap := cous.GetMappedByISO()

	var locs []*model.Location
	for _, locRow := range locsRow {
		loc, er := srv.defineLocation(locRow, &cousMap)
		if er.IsError() {
			return fltCount, addCount, er
		}
		locs = append(locs, &loc)
		addCount++

		if len(locs) >= saveBunch {
			if err := srv.locRps.CreateLocationsLocationsBulk(locs); err != nil {
				return fltCount, addCount, model.NewErrorServer("Error saving locations to database").WithError(err)
			}
			locs = []*model.Location{}
		}
	}
	if len(locs) > 0 {
		if err := srv.locRps.CreateLocationsLocationsBulk(locs); err != nil {
			return fltCount, addCount, model.NewErrorServer("Error saving locations to database").WithError(err)
		}
	}
	return fltCount, addCount, model.NewErrorNil()
}

// defineLocation defines location details
func (srv locationService) defineLocation(locRow model.LocationRow, cousMap *map[string]model.Country) (model.Location, model.Error) {
	var cou model.Country
	couLo := cou.GetCountryFromMapByISO(*cousMap, locRow.CountryCode)
	if couLo.ID == 0 {
		couNew, er := srv.couSrv.CreateCountry(locRow.CountryCode, locRow.Country)
		if er.IsError() {
			return model.Location{}, er
		}
		couLo = couNew
	}
	(*cousMap)[couLo.ISO] = couLo

	loc := model.DefineLocation(locRow.IPAddress, locRow.Latitude, locRow.Longitude, couLo.ID, locRow.City)
	return loc, model.NewErrorNil()
}

// sanitizeLocationsData sanitizes data filtering not valid data
func (srv locationService) sanitizeLocationsData(locs model.LocationRows) model.LocationRows {
	var rst model.LocationRows
	for _, loc := range locs {
		if !loc.IsAllRequiredData() {
			continue
		}

		ip := net.ParseIP(loc.IPAddress)
		if ip == nil {
			continue
		}

		lat, err := strconv.ParseFloat(loc.LatitudeStr, 64)
		if err != nil {
			continue
		}

		lon, err := strconv.ParseFloat(loc.LongitudeStr, 64)
		if err != nil {
			continue
		}

		loc.Latitude = lat
		loc.Longitude = lon
		rst = append(rst, loc)
	}
	return rst
}
