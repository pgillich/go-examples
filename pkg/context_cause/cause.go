package ctx_cause

import (
	"context"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	retryPeriod = 1 * time.Second
)

func RetryWithTimeout(ctx context.Context,
	f func() error, shutdownSignal chan struct{}, timeout time.Duration,
) error {
	ctx, cancel := context.WithCancelCause(ctx)
	doRetryEnd := make(chan struct{})
	defer close(doRetryEnd)

	go func() {
		select {
		case <-shutdownSignal:
			cancel(errors.New("cancel by shutdown"))
		case <-time.After(timeout):
			cancel(errors.New("cancel by timeout"))
		case <-doRetryEnd: // normal return
		}
	}()

	return RetryPeriod(ctx, f, retryPeriod)
}

func RetryPeriod(ctx context.Context,
	f func() error, period time.Duration,
) (err error) {
	ticker := time.NewTicker(period)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.WithError(context.Cause(ctx)).Warn("canceled")
			return err
		case <-ticker.C:
			if err = f(); err == nil {
				return err
			}
		}
	}
}
