package mq

import (
	"context"
	"github.com/adjust/rmq/v5"
	"github.com/semenovem/portal/pkg"
)

type MQ struct {
	ctx    context.Context
	logger pkg.Logger
	conn   rmq.Connection
}

type Conf struct {
	Tag          string
	Address      string
	DBNum        int
	GracefulExit func()
	ErrChan      chan<- error
}

func New(ctx context.Context, c *Conf, logger pkg.Logger) (*MQ, error) {
	o := MQ{
		logger: logger,
	}

	connection, err := rmq.OpenConnection(
		c.Tag,
		"tcp",
		c.Address,
		c.DBNum,
		c.ErrChan,
	)
	if err != nil {
		return nil, err
	}

	o.conn = connection

	go func() {
		<-ctx.Done()
		<-connection.StopAllConsuming()
		c.GracefulExit()
	}()

	return &o, nil
}

func (mq *MQ) OpenQueue(name string) (*Queue, error) {
	task, err := mq.conn.OpenQueue(name)
	if err != nil {
		return nil, err
	}

	q := &Queue{
		queue: task,
	}

	return q, nil
}
