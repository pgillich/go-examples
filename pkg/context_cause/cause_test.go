package ctx_cause

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	runTimeout = 10 * time.Second
)

func ping(wait time.Duration, err error) error {
	<-time.After(wait)
	return err
}

var ErrPingError = errors.New("ping error")

func TestOK(t *testing.T) {
	ping := func() error {
		return ping(0, nil)
	}
	shutdownSignal := make(chan struct{})

	err := RetryWithTimeout(context.Background(), ping, shutdownSignal, runTimeout)
	assert.NoError(t, err)
}

func TestErr(t *testing.T) {
	ping := func() error {
		return ping(0, ErrPingError)
	}
	shutdownSignal := make(chan struct{})

	err := RetryWithTimeout(context.Background(), ping, shutdownSignal, runTimeout)
	assert.Error(t, err)
	// time="2023-04-29T19:35:11+02:00" level=warning msg=canceled error="cancel by timeout"
}

func TestShutdown(t *testing.T) {
	ping := func() error {
		return ping(0, ErrPingError)
	}
	shutdownSignal := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second)
		close(shutdownSignal)
	}()

	err := RetryWithTimeout(context.Background(), ping, shutdownSignal, runTimeout)
	assert.Error(t, err)
	// time="2023-04-29T19:35:38+02:00" level=warning msg=canceled error="cancel by shutdown"
}
