package tomoe

import (
	"io"
	"net/http"
	"time"
)

// Client wraps the standard http.Client and provides helper methods.
type Client struct {
	baseURL    string
	httpClient *http.Client
	headers    map[string]string
	retries    int
	backoff    time.Duration
}

// RequestOptions defines options for the HTTP request.
type RequestOptions struct {
	Method      string
	Path        string
	QueryParams map[string]string
	Body        io.Reader
}
