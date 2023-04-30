package atomic_pointer

import (
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRequestDo(t *testing.T) {
	// given
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	log := logger.WithFields(logrus.Fields{"!TestCase": t.Name()})
	cleanupHttpDebug := RegisterHttpTracer(log)
	defer cleanupHttpDebug()
	testURL := "http://google.com"
	// when
	resp, err := ClientDo(http.MethodGet, testURL)
	// then
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
