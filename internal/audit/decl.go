package audit

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/proto/audit_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	Create    Action = "create"
	Delete    Action = "delete"
	Update    Action = "update"
	actionGet Action = "get"
)

const (
	Allow Decision = "allow" // Действие выполнено
	Deny  Decision = "deny"  // Запрет на действие
)

type Action string   // Типы действий в аудите
type Decision string // Решение на запрос
type Cause string    // Причина если было Deny

type P map[string]interface{}

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
	code    Code
	cause   Cause
	action  Action
	payload map[string]interface{}
}

func New(
	ctx context.Context,
	db *pgxpool.Pool,
	logger pkg.Logger,
	grpcConfig *config.GrpcClient,
) *AuditProvider {
	o := &AuditProvider{
		ctx:        ctx,
		logger:     logger.Named("AuditProvider"),
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
