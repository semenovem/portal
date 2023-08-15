package auditapp

import (
	"context"
	"fmt"
	"github.com/semenovem/portal/pkg"
	"google.golang.org/grpc/grpclog"
	"net"

	"google.golang.org/grpc"

	"github.com/semenovem/portal/proto/audit"
)

type GRPCServer struct {
	audit.AuditServer
	logger pkg.Logger
	ctx    context.Context
}

func (a *GRPCServer) RawString(ctx context.Context, request *audit.RawRequest) (*audit.RawResponse, error) {

	return nil, nil
}

func CreateGRPCServer() {
	var opts = make([]grpc.ServerOption, 0)

	grpcServer := grpc.NewServer(opts...)

	fmt.Println(">>>>>>> ", grpcServer)

	audit.RegisterAuditServer(grpcServer, &GRPCServer{
		AuditServer: nil,
	})

	listener, err := net.Listen("tcp", fmt.Sprint(":", 9090))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	if err = grpcServer.Serve(listener); err != nil {
		grpclog.Fatalf("%s grpc serve err: %v", "", err)
	}

	fmt.Println(">>>>>>>>>>>>> err = ", err)
}
