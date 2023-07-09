package signalmultipleclose

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMultipleClose_Panic(t *testing.T) {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{ //nolint:exhaustruct // optional members
		TimestampFormat: time.RFC3339Nano,
	}
	log := logger.WithFields(logrus.Fields{"!TestCase": t.Name()})

	var signalHandler Signal //nolint:gosimple // force interface
	signalHandler = NewMultipleClosePanic()

	exit1 := make(chan struct{})
	exit2 := make(chan struct{})

	go func(sh Signal) {
		log.Info("Closing #1")
		sh.Close()
		log.Info("Closed #1")
		close(exit1)
	}(signalHandler)

	<-exit1
	<-signalHandler.Done()

	assert.Equal(t, 1, signalHandler.CloseCount())

	go func(sh Signal) {
		assert.Panics(t, func() {
			log.Info("Closing #2")
			sh.Close()
			log.Info("Closed #2, not paniced!")
		})
		close(exit2)
	}(signalHandler)

	<-exit2
	<-signalHandler.Done()

	assert.Equal(t, 2, signalHandler.CloseCount())
}

func TestMultipleClose(t *testing.T) {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{ //nolint:exhaustruct // optional members
		TimestampFormat: time.RFC3339Nano,
	}
	log := logger.WithFields(logrus.Fields{"!TestCase": t.Name()})

	grNum := 2
	sleepMicro := 200
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	sleeps := make([]time.Duration, grNum)
	for gr := range sleeps {
		sleeps[gr] = time.Duration(random.Intn(sleepMicro)) * time.Microsecond
	}
	log.Infof("Sleeps: %v", sleeps)

	tests := []struct {
		name          string
		signalHandler Signal
	}{
		{"MultipleCloseNotsafe", NewMultipleCloseNotsafe()},
		{"NewMultipleCloseRecover", NewMultipleCloseRecover(log)},
		{"NewMultipleCloseOnce", NewMultipleCloseOnce()},
		{"MultipleCloseContext", NewMultipleCloseContext()},

		// Go 1.21
		// {"NewMultipleCloseOncefunc", NewMultipleCloseOncefunc()},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			log := log.WithFields(logrus.Fields{"!TestCase": t.Name()})

			wgBegin := sync.WaitGroup{}
			wgBegin.Add(grNum)
			wgEnd := sync.WaitGroup{}
			wgEnd.Add(grNum)
			doStart := make(chan struct{})

			for gr := 0; gr < grNum; gr++ {
				go func(sh Signal, gr int) {
					time.Sleep(sleeps[gr])
					wgBegin.Done()
					<-doStart
					log.Infof("Closing #%d", gr)
					sh.Close()
					log.Infof("Closed #%d", gr)
					wgEnd.Done()
				}(tt.signalHandler, gr)
			}

			wgBegin.Wait()
			close(doStart)
			wgEnd.Wait()
			tt.signalHandler.Done()

			assert.Equal(t, grNum, tt.signalHandler.CloseCount())
		})
	}
}
