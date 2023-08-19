package grpc_server

import (
	"context"
	"fmt"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/proto/audit_grpc"
	"google.golang.org/grpc"
	"net"
	"time"
)

type GRPCServer struct {
	logger pkg.Logger
	ctx    context.Context
	audit_grpc.AuditServer
	maxProcessingTime time.Duration
}

func New(
	ctx context.Context,
	logger pkg.Logger,
	cfg *config.GrpcServer,
) (*GRPCServer, error) {
	var (
		ll         = logger.Named("GRPCServer.New")
		opts       = make([]grpc.ServerOption, 0)
		grpcServer = grpc.NewServer(opts...)
		impl       = &GRPCServer{
			ctx:               ctx,
			logger:            logger.Named("GRPCServer"),
			maxProcessingTime: cfg.GetMaxProcessingTimeSec(),
		}
	)

	fmt.Println(">>>>>>> ", grpcServer)

	audit_grpc.RegisterAuditServer(grpcServer, impl)

	listener, err := net.Listen("tcp", fmt.Sprint(":", cfg.Port))
	if err != nil {
		ll.Named("Listen").Error(err.Error())
		return nil, err
	}

	go func() {
		if err = grpcServer.Serve(listener); err != nil {
			ll.Named("Serve").Error(err.Error())
		}
	}()

	fmt.Println(">>>>>>>>>>>>> err = ", err)

	return impl, nil
}
