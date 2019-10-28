package sql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gitlab.com/dpcat237/geolocation/config"
	"gitlab.com/dpcat237/geolocation/model"
	"strings"
)

// DatabaseCollector defines required methods for collector of database connections
type DatabaseCollector interface {
	GetDatabase() *gorm.DB
	GracefulStop()
}

// databaseCollector defines database connections
type databaseCollector struct {
	Db *gorm.DB
}

// CheckNotFound checks if database error is referred to not found data
func CheckNotFound(err error) bool {
	return gorm.IsRecordNotFoundError(err) || (err != nil && err.Error() == "sql: no rows in result set")
}

// GetDatabase returns database connection
func (cll *databaseCollector) GetDatabase() *gorm.DB {
	return cll.Db
}

// GracefulStop stops database connections
func (cll *databaseCollector) GracefulStop() {
	cll.Db.Close()
}

// InitDbCollector initializes database connections and set to collector
func InitDbCollector(dbDSN string, mode string) (DatabaseCollector, model.Error) {
	var cll databaseCollector
	db, er := initConnectionDb(dbDSN, mode)
	if er.IsError() {
		return &cll, er
	}
	cll.Db = db
	return &cll, model.NewErrorNil()
}

// bulkCreate define bunch query and lunch it to create data on database
func bulkCreate(cn *gorm.DB, tabName, insFields string, valStrs []string, valArgs []interface{}) error {
	smt := `INSERT INTO ` + tabName + `(` + insFields + `) VALUES %s`
	smt = fmt.Sprintf(smt, strings.Join(valStrs, ","))
	tx := cn.Begin()
	if err := tx.Exec(smt, valArgs...).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// initConnectionDb initializes database connection
func initConnectionDb(dbDSN, mode string) (*gorm.DB, model.Error) {
	db, err := gorm.Open("mysql", dbDSN+"&parseTime=true")
	if err != nil {
		return db, model.NewErrorServer("Error opening database connection").WithError(err)
	}
	if mode != config.ModeProd {
		db.LogMode(true)
	}
	return db, model.NewErrorNil()
}
