package router

import (
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"gitlab.com/dpcat237/geolocation/controller"
	"gitlab.com/dpcat237/geolocation/logger"
	"gitlab.com/dpcat237/geomicroservices/geolocation"
)

// Manager defines required methods for router manager
type Manager interface {
	Init()
	GracefulStop()
}

// manager defines required services for router manager
type manager struct {
	logg     logger.Logger
	grpc     *grpc.Server
	grpcPort string
}

// New initializes router manager
func New(logg logger.Logger, ctrCll controller.Collector, pGr string) *manager {
	return &manager{
		logg:     logg,
		grpc:     createGRPC(ctrCll),
		grpcPort: pGr,
	}
}

// Init runs gRPC connection
func (mng *manager) Init() {
	con, err := net.Listen("tcp", ":"+mng.grpcPort)
	if err != nil {
		mng.logg.WithError(err).Fatal("Failed to start gRPC connection")
	}
	mng.logg.Infof("Connection for gRPC created at %s on port %s", time.Now().String(), mng.grpcPort)
	go func() {
		if err := mng.grpc.Serve(con); err != nil {
			mng.logg.WithError(err).Fatal("Failed to start gRPC")
		}
	}()
}

// GracefulStop stops gRPC connection
func (mng *manager) GracefulStop() {
	mng.grpc.GracefulStop()
	mng.logg.Info("Stopped location gRPC")
}

// createGRPC creates gRPC server
func createGRPC(ctrCll controller.Collector) *grpc.Server {
	gr := grpc.NewServer()
	geolocation.RegisterLocationControllerServer(gr, ctrCll.LocCtr)
	reflection.Register(gr)
	return gr
}
