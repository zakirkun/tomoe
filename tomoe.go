package tomoe

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// NewClient initializes a new HTTP client with a base URL and timeout.
func NewClient(baseURL string, timeout time.Duration, retries int, backoff time.Duration, headers map[string]string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
		headers: headers,
		retries: retries,
		backoff: backoff,
	}
}

func (c *Client) ParallelRequests(ctx context.Context, opts []RequestOptions) ([]*http.Response, error) {
	results := make([]*http.Response, len(opts))
	errCh := make(chan error, len(opts))
	doneCh := make(chan struct{}, len(opts))

	for i, opt := range opts {
		go func(i int, opt RequestOptions) {
			defer func() { doneCh <- struct{}{} }()
			res, err := c.Do(ctx, opt)
			if err != nil {
				errCh <- err
				return
			}
			results[i] = res
		}(i, opt)
	}

	// Wait for all requests to finish
	for i := 0; i < len(opts); i++ {
		<-doneCh
	}

	// Check for errors
	close(errCh)
	if len(errCh) > 0 {
		return nil, <-errCh
	}

	return results, nil
}

func (c *Client) Do(ctx context.Context, opts RequestOptions) (*http.Response, error) {
	var lastErr error

	for attempt := 1; attempt <= c.retries; attempt++ {
		if attempt > 1 {
			time.Sleep(c.backoff) // Exponential backoff
		}

		result, err := c.executeRequest(ctx, opts)
		if err != nil {
			lastErr = fmt.Errorf("attempt %d failed: %w", attempt, err)
			continue
		}

		return result, nil
	}

	return nil, fmt.Errorf("all retries failed: %v", lastErr)
}

func (c *Client) executeRequest(ctx context.Context, opts RequestOptions) (*http.Response, error) {

	// Construct the URL
	reqURL, err := url.Parse(c.baseURL + opts.Path)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	// Add query parameters
	q := reqURL.Query()
	for key, value := range opts.QueryParams {
		q.Add(key, value)
	}
	reqURL.RawQuery = q.Encode()

	// Create the request
	req, err := http.NewRequestWithContext(ctx, opts.Method, reqURL.String(), opts.Body)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	// Perform the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}
