package throwgrpc

import (
	"fmt"
	"github.com/semenovem/portal/pkg/throw"
	"github.com/semenovem/portal/pkg/throw/throwtrace"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	reasonDomain = ""
)

func SetReasonDomain(s string) {
	reasonDomain = s
}

func From(err error) *throw.Throw {
	if s, ok := status.FromError(err); ok {
		var th *throw.Throw

		switch s.Code() {
		case codes.NotFound:
			th = throw.NewNotFound(err)
		case codes.InvalidArgument:
			th = throw.NewBadRequest(err)
		case codes.PermissionDenied:
			th = throw.NewDeny(err)
		case codes.Unauthenticated:
			th = throw.NewDeny(err)
		case codes.AlreadyExists:
			th = throw.NewDuplicate(err)
		}

		if st := s.Proto(); st != nil {
			for _, v := range st.Details {
				fmt.Println(">>>>>>>> v = ", string(v.Value))
			}
		}

		return th
	}

	return throw.Cast(err)
}

func To(err error) error {
	var (
		th = throw.Cast(err)
		st *status.Status
	)

	switch th.Kind() {
	case throw.BadRequest:
		st = status.New(codes.InvalidArgument, th.Error())
	case throw.NotFound:
		st = status.New(codes.NotFound, th.Error())
	case throw.Deny:
		st = status.New(codes.PermissionDenied, th.Error())
	case throw.Auth:
		st = status.New(codes.Unauthenticated, th.Error())
	case throw.Duplicate:
		st = status.New(codes.AlreadyExists, th.Error())
	default:
		st = status.New(codes.Internal, th.Error())
	}

	for _, r := range throwtrace.Reasons(th) {
		e := &errdetails.ErrorInfo{
			Reason:   r.Reason,
			Domain:   reasonDomain,
			Metadata: r.Metadata,
		}

		if st, err = st.WithDetails(e); err != nil {
			return status.Error(codes.Internal, "error while preparing grpc response: "+err.Error())
		}
	}

	return st.Err()
}
