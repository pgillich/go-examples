package atomic_pointer

import (
	"context"
	"net/http"
	"sync/atomic"
)

// defaultHttpClientTransport is used as http.Client.Transport
// for next outgoing requests.
var defaultHttpClientTransport atomic.Pointer[http.RoundTripper]

// init sets the default http.Transport
func init() {
	defaultHttpClientTransport.Store(&http.DefaultTransport)
}

// ClientDo uses defaultHttpClientTransport to make outgoing http requests
func ClientDo(method, url string) (*http.Response, error) {
	req, _ := http.NewRequestWithContext(context.Background(), method, url, nil)

	client := http.Client{Transport: *defaultHttpClientTransport.Load()}

	return client.Do(req)
}
