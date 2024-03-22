package channel

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	// "github.com/stretchr/testify/assert"
)

func TestWriteToLaterClosed(t *testing.T) {
	// given
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	log := logger.WithFields(logrus.Fields{"!TestCase": t.Name()})
	targetCh := make(chan int)
	started := make(chan struct{})
	go func() {
		close(started)
		log.Infof("Started")
		targetCh <- 1
		log.Infof("End")
	}()
	<-started
	time.Sleep(1 * time.Second)
	log.Infof("Begin")

	// when
	close(targetCh)

	// then
	log.Infof("Finish")
}

func TestWriteToClosedBefore(t *testing.T) {
	// given
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	log := logger.WithFields(logrus.Fields{"!TestCase": t.Name()})
	targetCh := make(chan int)
	close(targetCh)

	// when
	log.Infof("Started")
	targetCh <- 1
	log.Infof("End")

	// then
	log.Infof("Finish")
}

func TestWriteToClosedBeforeSelect(t *testing.T) {
	// given
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	log := logger.WithFields(logrus.Fields{"!TestCase": t.Name()})
	targetCh := make(chan int)
	close(targetCh)

	// when
	log.Infof("Started")
	select {
	case targetCh <- 1:
		log.Infof("Written")
	default:
	}
	log.Infof("End")

	// then
	log.Infof("Finish")
}

func TestCloseNil(t *testing.T) {
	// given
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	log := logger.WithFields(logrus.Fields{"!TestCase": t.Name()})
	var targetCh chan int

	// when
	log.Infof("Closing")
	close(targetCh)

	// then
	log.Infof("Finish")
}

func TestWriteToNil(t *testing.T) {
	// given
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	log := logger.WithFields(logrus.Fields{"!TestCase": t.Name()})
	var targetCh chan int

	// when
	log.Infof("Started")
	select {
	case targetCh <- 1:
		log.Infof("Written")
	default:
		log.Infof("Blocked")
	}
	log.Infof("End")

	// then
	log.Infof("Finish")
}

func TestReadNil(t *testing.T) {
	// given
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	log := logger.WithFields(logrus.Fields{"!TestCase": t.Name()})
	var targetCh chan int

	// when
	log.Infof("Started")
	select {
	case <-targetCh:
		log.Infof("Read")
	default:
		log.Infof("Blocked")
	}
	log.Infof("End")

	// then
	log.Infof("Finish")
}

func TestWriteToLaterNil(t *testing.T) {
	// given
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	log := logger.WithFields(logrus.Fields{"!TestCase": t.Name()})
	targetCh := make(chan int)
	started := make(chan struct{})
	go func() {
		close(started)
		log.Infof("Started")
		targetCh <- 1
		log.Infof("End")
	}()
	<-started
	time.Sleep(1 * time.Second)
	log.Infof("Begin")

	// when
	targetCh = nil
	time.Sleep(1 * time.Second)

	// then
	log.Infof("Finish")
}

func TestWriteToLaterNilP(t *testing.T) {
	// given
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	log := logger.WithFields(logrus.Fields{"!TestCase": t.Name()})
	targetCh := make(chan int)
	targetChP := &targetCh
	var targetChNil chan int
	started := make(chan struct{})
	go func() {
		close(started)
		log.Infof("Started")
		targetCh <- 1
		log.Infof("End")
	}()
	<-started
	time.Sleep(1 * time.Second)
	log.Infof("Begin")

	// when
	*targetChP = targetChNil
	time.Sleep(1 * time.Second)

	// then
	log.Infof("Finish")
}
