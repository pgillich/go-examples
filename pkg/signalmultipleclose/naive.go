package signalmultipleclose

import (
	"sync/atomic"
)

type MultipleCloseNotsafe struct {
	signal     chan struct{}
	closeCount atomic.Int32
}

func NewMultipleCloseNotsafe() *MultipleCloseNotsafe {
	return &MultipleCloseNotsafe{ //nolint:exhaustruct // closeCount zero value is good
		signal: make(chan struct{}),
	}
}

func (s *MultipleCloseNotsafe) Close() {
	s.closeCount.Add(1)

	select {
	case <-s.signal:
	default:
		close(s.signal)
	}
}

func (s *MultipleCloseNotsafe) Done() <-chan struct{} {
	return s.signal
}

func (s *MultipleCloseNotsafe) CloseCount() int {
	return int(s.closeCount.Load())
}
