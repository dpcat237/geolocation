package sql

import (
	"github.com/jinzhu/gorm"

	"gitlab.com/dpcat237/geolocation/model"
)

// CountryRepository defines required methods for Country repository
type CountryRepository interface {
	GetCountries() (model.Countries, error)
	SaveCountry(c *model.Country) error
}

// countryRepository defines required services for Country repository
type countryRepository struct {
	db *gorm.DB
}

// NewCountry initializes Country repository
func NewCountry(db *gorm.DB) *countryRepository {
	return &countryRepository{db: db}
}

// GetCountries returns countries from database
func (rps countryRepository) GetCountries() (model.Countries, error) {
	var cous []model.Country
	return cous, rps.db.Find(&cous).Error
}

// SaveCountry creates / updates country details on database
func (rps countryRepository) SaveCountry(c *model.Country) error {
	return rps.db.Save(&c).Error
}
