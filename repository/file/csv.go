package file

import (
	"encoding/csv"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/gocarina/gocsv"

	"gitlab.com/dpcat237/geolocation/model"
)

const (
	// columnSeparator defines column separator for CSV file
	columnSeparator = ','
	// fileTmpFolderPath defines location of CSV files
	fileTmpFolderPath = "repository/file/data/"
)

// CSVClient defines required methods for CSV client
type CSVClient interface {
	GetLocations() (model.LocationRows, model.Error)
	GracefulClose() model.Error
}

// csvClient defines required services for CSV client
type csvClient struct {
	cltFl *os.File
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewClient initializes CSV client
func NewClient(flName string) (CSVClient, model.Error) {
	cli := csvClient{cltFl: &os.File{}}
	if flName == "" {
		return &cli, model.NewErrorNil()
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = columnSeparator
		return r
	})

	cltFl, err := os.Open(fileTmpFolderPath + flName)
	if err != nil {
		return &cli, model.NewErrorServer("Error to open a file " + flName).WithError(err)
	}
	cli.cltFl = cltFl
	return &cli, model.NewErrorNil()
}

// GetLocations extracts location rows from CSV file
func (mng *csvClient) GetLocations() (model.LocationRows, model.Error) {
	var rows model.LocationRows
	if err := gocsv.UnmarshalFile(mng.cltFl, &rows); err != nil {
		return nil, model.NewErrorServer("Error parsing file").WithError(err)
	}
	return rows, model.NewErrorNil()
}

// GracefulClose stops CSV client
func (mng *csvClient) GracefulClose() model.Error {
	if err := mng.cltFl.Close(); err != nil {
		return model.NewErrorServer("Error closing client").WithError(err)
	}
	return model.NewErrorNil()
}
