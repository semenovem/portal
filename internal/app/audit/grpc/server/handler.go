package server

import (
	"context"
	"fmt"
	"github.com/semenovem/portal/proto/audit"
)

func (a *GRPCServer) RawString(ctx context.Context, request *audit.RawRequest) (*audit.RawResponse, error) {

	fmt.Printf(">>> %s. %s\n", request.Payload, request.ID)

	response := &audit.RawResponse{
		Ok: 2,
	}

	return response, nil
}
