package signalmultipleclose

import (
	"sync/atomic"
)

type MultipleClosePanic struct {
	signal     chan struct{}
	closeCount atomic.Int32
}

func NewMultipleClosePanic() *MultipleClosePanic {
	return &MultipleClosePanic{ //nolint:exhaustruct // closeCount zero value is good
		signal: make(chan struct{}),
	}
}

func (s *MultipleClosePanic) Close() {
	s.closeCount.Add(1)

	close(s.signal)
}

func (s *MultipleClosePanic) Done() <-chan struct{} {
	return s.signal
}

func (s *MultipleClosePanic) CloseCount() int {
	return int(s.closeCount.Load())
}
