package service

import (
	"gitlab.com/dpcat237/geolocation/model"
	"gitlab.com/dpcat237/geolocation/repository/sql"
)

// CountryService defines required methods for Country service
type CountryService interface {
	CreateCountry(iso, nm string) (model.Country, model.Error)
	GetCountries() (model.Countries, model.Error)
}

// countryService defines required services for Country service
type countryService struct {
	couRps sql.CountryRepository
}

// newCountry initializes Country service
func newCountry(couRps sql.CountryRepository) *countryService {
	return &countryService{couRps: couRps}
}

// CreateCountry creates country on database
func (srv countryService) CreateCountry(iso, nm string) (model.Country, model.Error) {
	cou := model.Country{
		ISO:  iso,
		Name: nm,
	}

	if err := srv.couRps.SaveCountry(&cou); err != nil {
		return cou, model.NewErrorServer("Error saving country").WithError(err)
	}
	return cou, model.NewErrorNil()
}

// GetCountries returns countries from database
func (srv countryService) GetCountries() (model.Countries, model.Error) {
	cous, err := srv.couRps.GetCountries()
	if err != nil && !sql.CheckNotFound(err) {
		return cous, model.NewErrorServer("Error getting countries").WithError(err)
	}
	return cous, model.NewErrorNil()
}
