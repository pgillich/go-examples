package atomic_pointer

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// HttpTracer is a client middleware to log requests
type HttpTracer struct {
	orig *http.RoundTripper
	log  *logrus.Entry
}

// RoundTrip implements http.RoundTripper, logs debug info.
func (tr *HttpTracer) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := (*tr.orig).RoundTrip(req)

	fields := logrus.Fields{
		"ReqMethod":  req.Method,
		"ReqURL":     req.URL,
		"ReqHeaders": req.Header,
	}
	if resp != nil {
		fields["RespStatusCode"] = resp.StatusCode
		fields["RespHeaders"] = resp.Header
	}
	tr.log.WithFields(fields).Debug("Client request")

	return resp, err
}

// RegisterHttpTracer registers HttpTracer as a client middleware.
// Returns the deregister function.
func RegisterHttpTracer(log *logrus.Entry) func() {
	httpTracer := &HttpTracer{defaultHttpClientTransport.Load(), log}
	var rt http.RoundTripper = httpTracer

	defaultHttpClientTransport.Store(&rt)

	return func() {
		defaultHttpClientTransport.Store(httpTracer.orig)
	}
}
