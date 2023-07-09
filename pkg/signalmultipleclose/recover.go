package signalmultipleclose

import (
	"sync/atomic"

	"github.com/sirupsen/logrus"
)

type MultipleCloseRecover struct {
	signal     chan struct{}
	closeCount atomic.Int32
	log        *logrus.Entry
}

func NewMultipleCloseRecover(log *logrus.Entry) *MultipleCloseRecover {
	return &MultipleCloseRecover{ //nolint:exhaustruct // closeCount zero value is good
		signal: make(chan struct{}),
		log:    log,
	}
}

func (s *MultipleCloseRecover) Close() {
	s.closeCount.Add(1)

	defer func() {
		if recover() != nil {
			s.log.Info("Closed with panic")
		}
	}()

	close(s.signal)
}

func (s *MultipleCloseRecover) Done() <-chan struct{} {
	return s.signal
}

func (s *MultipleCloseRecover) CloseCount() int {
	return int(s.closeCount.Load())
}
