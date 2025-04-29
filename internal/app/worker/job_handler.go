package worker

import "context"


type JobHandler interface {
	Handle(ctx context.Context, job Job) error
	SetNext(next JobHandler)
}

type BaseHandler struct {
	next JobHandler
}

func (b *BaseHandler) Next(ctx context.Context, job Job) error {
	if b.next != nil {
		return b.next.Handle(ctx, job)
	}
	return nil
}

func (b *BaseHandler) SetNext(next JobHandler) {
	b.next = next
}
