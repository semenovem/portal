package config

import "time"

type Audit struct {
	Base

	Grpc struct {
		ServerPort           string `env:"GRPC_SERVER_PORT,required"`
		MaxProcessingTimeSec uint32 `env:"GRPC_MAX_PROCESSING_TIME_SEC" envDefault:"5"`
	}
}

func (a *Audit) GetGrpcMaxProcessingTimeSec() time.Duration {
	return time.Second * time.Duration(a.Grpc.MaxProcessingTimeSec)
}
