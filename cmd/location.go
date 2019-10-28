package cmd

import (
	"time"

	"github.com/spf13/cobra"

	"gitlab.com/dpcat237/geolocation/config"
	"gitlab.com/dpcat237/geolocation/logger"
	"gitlab.com/dpcat237/geolocation/repository/sql"
	"gitlab.com/dpcat237/geolocation/service"
)

// grpcCmd defines CLI command to lunch data loader process
var locationsCmd = &cobra.Command{
	Use:   "load-locations",
	Short: "Load locations from CSV",
	Run: func(cmd *cobra.Command, args []string) {
		loadUsersLocations()
	},
}

func init() {
	rootCmd.AddCommand(locationsCmd)
}

// loadUsersLocations runs process to load locations data from CSV file
func loadUsersLocations() {
	cfg := config.LoadConfigData()
	log := logger.New()
	dbCl, er := sql.InitDbCollector(cfg.DbDSN, cfg.Mode)
	if er.IsError() {
		log.Error(er)
	}

	srvCll := service.Init(dbCl)
	startAt := time.Now()
	fltCount, addCount, er := srvCll.LocSrv.LoadLocations()
	if er.IsError() {
		log.Println("Error loading locations data", er.GetLogMessage())
		return
	}
	log.Printf("Successfully loaded %d locations and filtered %d. Process took %s", addCount, fltCount, time.Since(startAt))
}
