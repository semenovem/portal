package audit_provider

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/audit"
	"github.com/semenovem/portal/proto/audit_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	userID  uint32
	code    audit.Code
	cause   audit.Cause
	payload map[string]interface{}
}

func NewAudit(
	ctx context.Context,
	db *pgxpool.Pool,
	logger pkg.Logger,
	grpcConfig *config.GrpcClient,
) *AuditProvider {
	o := &AuditProvider{
		ctx:        ctx,
		logger:     logger.Named("authPvd"),
		db:         db,
		grpcConfig: grpcConfig,
		input:      make(chan *auditPipe, 1000),
	}

	go o.processing()

	return o
}

func (a *AuditProvider) connectGrpc() error {
	ll := a.logger.Named("connectGrpc")

	conn, err := grpc.Dial(a.grpcConfig.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		ll.Error(err.Error())
		return err
	}

	a.conn = conn
	a.client = audit_grpc.NewAuditClient(conn)

	return nil
}
