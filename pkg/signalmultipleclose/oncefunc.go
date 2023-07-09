//go:build go1.21

package signalmultipleclose

import (
	"sync"
	"sync/atomic"
)

type MultipleCloseOncefunc struct {
	signal     chan struct{}
	closeCount atomic.Int32
	once       func()
}

func NewMultipleCloseOncefunc() *MultipleCloseOncefunc {
	signal := make(chan struct{})

	return &MultipleCloseOncefunc{ //nolint:exhaustruct // closeCount zero value is good
		signal: signal,
		once: sync.OnceFunc(func() { // Introduced in Go 1.21
			close(signal)
		}),
	}
}

func (s *MultipleCloseOncefunc) Close() {
	s.closeCount.Add(1)

	s.once()
}

func (s *MultipleCloseOncefunc) Done() <-chan struct{} {
	return s.signal
}

func (s *MultipleCloseOncefunc) CloseCount() int {
	return int(s.closeCount.Load())
}
