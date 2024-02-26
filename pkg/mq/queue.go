package mq

import "github.com/adjust/rmq/v5"

type Queue struct {
	queue rmq.Queue
}

func (q *Queue) Send(b []byte) error {
	return q.queue.PublishBytes(b)
}
