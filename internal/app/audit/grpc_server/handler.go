package grpc_server

import (
	"context"
	"fmt"
	"github.com/semenovem/portal/proto/audit_grpc"
)

func (a *GRPCServer) RawString(ctx context.Context, request *audit_grpc.RawRequest) (*audit_grpc.RawResponse, error) {

	fmt.Printf(">>> %s. %s\n", request.Payload, request.ID)

	response := &audit_grpc.RawResponse{
		Ok: 2,
	}

	return response, nil
}
