package signalmultipleclose

import (
	"context"
	"sync/atomic"
)

type MultipleCloseContext struct {
	ctx        context.Context //nolint:containedctx // local context
	cancel     func()
	closeCount atomic.Int32
}

func NewMultipleCloseContext() *MultipleCloseContext {
	ctx, cancel := context.WithCancel(context.Background())

	return &MultipleCloseContext{ //nolint:exhaustruct // closeCount zero value is good
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *MultipleCloseContext) Close() {
	s.closeCount.Add(1)

	s.cancel()
}

func (s *MultipleCloseContext) Done() <-chan struct{} {
	return s.ctx.Done()
}

func (s *MultipleCloseContext) CloseCount() int {
	return int(s.closeCount.Load())
}
