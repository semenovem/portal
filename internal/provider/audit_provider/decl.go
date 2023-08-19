package audit_provider

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/audit"
	"github.com/semenovem/portal/proto/audit_grpc"
	"google.golang.org/grpc"
)

type AuditProvider struct {
	ctx        context.Context
	db         *pgxpool.Pool
	logger     pkg.Logger
	grpcConfig *config.GrpcClient
	conn       *grpc.ClientConn
	client     audit_grpc.AuditClient
	input      chan *auditPipe
}

type auditPipe struct {
	code    audit.Code
	cause   audit.Cause
	payload map[string]interface{}
}
