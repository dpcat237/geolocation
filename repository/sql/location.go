package sql

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/dpcat237/geolocation/model"
)

// LocationRepository defines required methods for Location repository
type LocationRepository interface {
	GetLastLocationByIP(ip string) (model.Location, error)
	CreateLocationsLocationsBulk(los []*model.Location) error
}

// sessionRepository defines required services for Location repository
type sessionRepository struct {
	db *gorm.DB
}

// NewLocation initializes Location repository
func NewLocation(db *gorm.DB) *sessionRepository {
	return &sessionRepository{db: db}
}

// GetLastLocationByIP get last location details by IP address
func (rps sessionRepository) GetLastLocationByIP(ip string) (model.Location, error) {
	var loc model.Location
	q := rps.db.Where("ip_address = ?", ip).Order("created_at DESC").Preload("Country")
	return loc, q.First(&loc).Error
}

// CreateLocationsLocationsBulk creates locations in a bunch on database
func (rps sessionRepository) CreateLocationsLocationsBulk(los []*model.Location) error {
	var valStrings []string
	var valArgs []interface{}
	columns := `ip_address, latitude, longitude, country_id, city, created_at`

	for _, lo := range los {
		valStrings = append(valStrings, "(?, ?, ?, ?, ?, ?)")

		valArgs = append(valArgs, lo.IPAddress)
		valArgs = append(valArgs, lo.Latitude)
		valArgs = append(valArgs, lo.Longitude)
		valArgs = append(valArgs, lo.CountryID)
		valArgs = append(valArgs, lo.City)
		valArgs = append(valArgs, lo.CreatedAt)
	}
	return bulkCreate(rps.db, model.LocationTable, columns, valStrings, valArgs)
}
