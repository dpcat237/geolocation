package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/dpcat237/geolocation/config"
	"gitlab.com/dpcat237/geolocation/controller"
	"gitlab.com/dpcat237/geolocation/logger"
	"gitlab.com/dpcat237/geolocation/repository/sql"
	"gitlab.com/dpcat237/geolocation/router"
	"gitlab.com/dpcat237/geolocation/service"
)

// grpcCmd defines CLI command to lunch gRPC
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start GRPC",
	Long:  `Start GRPC to process requests`,
	Run: func(cmd *cobra.Command, args []string) {
		startGRPC()
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}

// startGRPC initializes services to exchange data via gRPC
func startGRPC() {
	cfg := config.LoadConfigData()
	logg := logger.New()
	dbCl, er := sql.InitDbCollector(cfg.DbDSN, cfg.Mode)
	if er.IsError() {
		logg.Error(er)
	}

	srvCll := service.Init(dbCl)
	ctrCll := controller.InitCollector(logg, srvCll)
	rtrMng := router.New(logg, ctrCll, cfg.GRPCport)
	rtrMng.Init()

	gracefulStop(logg, dbCl, rtrMng, ctrCll)
	logg.Infof("Service stopped")
}

// gracefulStop stop gRPC and database connections
func gracefulStop(logg logger.Logger, dbCl sql.DatabaseCollector, rtrMng router.Manager, ctrCll controller.Collector) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c
	close(c)
	logg.Info("Closing...")
	dbCl.GracefulStop()
	rtrMng.GracefulStop()
}
