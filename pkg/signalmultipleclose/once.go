package signalmultipleclose

import (
	"sync"
	"sync/atomic"
)

type MultipleCloseOnce struct {
	signal     chan struct{}
	closeCount atomic.Int32
	once       sync.Once
}

func NewMultipleCloseOnce() *MultipleCloseOnce {
	return &MultipleCloseOnce{ //nolint:exhaustruct // closeCount zero value is good
		signal: make(chan struct{}),
	}
}

func (s *MultipleCloseOnce) Close() {
	s.closeCount.Add(1)

	s.once.Do(func() {
		close(s.signal)
	})
}

func (s *MultipleCloseOnce) Done() <-chan struct{} {
	return s.signal
}

func (s *MultipleCloseOnce) CloseCount() int {
	return int(s.closeCount.Load())
}
