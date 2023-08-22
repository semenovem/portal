package audit

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/semenovem/portal/proto/audit_grpc"
)

func (a *AuditProvider) processing() {
	ll := a.logger.Named("processing")

	ll.Info("started")

	for one := range a.input {
		var s string

		payload, err := json.Marshal(one.payload)
		if err != nil {
			ll.Named("Marshal").With("payload", one.payload).Error(err.Error())
			continue
		}

		if one.userID != 0 {
			s = fmt.Sprintf("[userID: %d]", one.userID)
		}

		s += fmt.Sprintf("[code:%s]", one.code)

		s += fmt.Sprintf("[payload:%s]", string(payload))

		//if one.cause == "" {
		//	s = fmt.Sprintf("Approved: code=%s  payload:[%+v]", one.code, string(payload))
		//} else {
		//	s = fmt.Sprintf("Refusal: code=%s  cause=%s   payload:[%+v]",
		//		one.code, one.cause, string(payload))
		//}

		//  todo дополнительно записать в БД в схему audit

		if a.client == nil {
			if err := a.connectGrpc(); err != nil {
				ll.Named("connectGrpc").Nested(err.Error())
				ll.Named("reserved.output").Debug(s)
				continue
			}
		}

		_, err = a.client.RawString(a.ctx, &audit_grpc.RawRequest{
			ID:      uuid.NewString(),
			Payload: s,
		})
		if err != nil {
			ll.Named("RawString").Error(err.Error())
			continue
		}
	}
}
