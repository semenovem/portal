package audit_provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/audit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	proto "github.com/semenovem/portal/proto/audit"
)

type AuditProvider struct {
	ctx        context.Context
	db         *pgxpool.Pool
	logger     pkg.Logger
	grpcConfig *config.GrpcClient
	conn       *grpc.ClientConn
	client     proto.AuditClient
	input      chan *auditPipe
}

type auditPipe struct {
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
	a.client = proto.NewAuditClient(conn)

	return nil
}

func (a *AuditProvider) processing() {
	var (
		ll = a.logger.Named("processing")
	)

	ll.Info("started")

	for one := range a.input {
		var s string

		payload, err := json.Marshal(one.payload)
		if err != nil {
			ll.Named("Marshal").With("payload", one.payload).Error(err.Error())
			continue
		}

		if one.cause == "" {
			s = fmt.Sprintf("Approved: code=%s  payload:[%+v]", one.code, string(payload))
		} else {
			s = fmt.Sprintf("Refusal: code=%s  cause=%s   payload:[%+v]",
				one.code, one.cause, string(payload))
		}

		//  todo дополнительно записать в БД в схему audit

		if a.client == nil {
			if err := a.connectGrpc(); err != nil {
				ll.Named("connectGrpc").Nested(err.Error())
				ll.Named("reserved.output").Debug(s)
				continue
			}
		}

		_, err = a.client.RawString(a.ctx, &proto.RawRequest{
			ID:      uuid.NewString(),
			Payload: s,
		})
		if err != nil {
			ll.Named("RawString").Error(err.Error())
			continue
		}

		//response.Ok
	}
}

// Refusal отказ
func (a *AuditProvider) Refusal(code audit.Code, cause audit.Cause, payload map[string]interface{}) {
	a.input <- &auditPipe{
		code:    code,
		cause:   cause,
		payload: payload,
	}
}

// Approved успешно
func (a *AuditProvider) Approved(code audit.Code, payload map[string]interface{}) {
	a.input <- &auditPipe{
		code:    code,
		payload: payload,
	}
}

// Send Аудит аутентификации
func (a *AuditProvider) Send(code audit.Code, decision audit.Decision, payload map[string]interface{}) {
	fmt.Println("audit >>>>>>>>>> code = ", code)
	fmt.Println("audit >>>>>>>>>> decision = ", decision)
	fmt.Println("audit >>>>>>>>>> payload = ", payload)
}

// User Аудит действий пользователя
func (a *AuditProvider) User(userID uint32, action audit.Action, payload map[string]interface{}) {
	fmt.Println("audit >>>>>>>>>> userID = ", userID)
	fmt.Println("audit >>>>>>>>>> decision = ", action)
	fmt.Println("audit >>>>>>>>>> payload = ", payload)
}

// Аудит действий пользователя
