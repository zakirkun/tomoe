package tomoe

import (
	"context"
	"io"
	"testing"
	"time"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func TestSingle(t *testing.T) {
	client := NewClient("https://jsonplaceholder.typicode.com", 30*time.Second, 3, 5*time.Second, nil)
	ctx := context.Background()

	// Single request with retries
	opts := RequestOptions{
		Method: "GET",
		Path:   "/todos/1",
	}

	response, err := client.Do(ctx, opts)
	if err != nil {
		t.Errorf("Single Request Error: %v", err.Error())
		return
	}
	defer response.Body.Close()

	// Read raw body for additional processing or error messages
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Parse Body Error: %v", err.Error())
		return
	}

	t.Logf("Success: %v", string(body))
}

func TestParrarel(t *testing.T) {

	client := NewClient("https://jsonplaceholder.typicode.com", 30*time.Second, 3, 5*time.Second, nil)
	ctx := context.Background()

	requests := []RequestOptions{
		{Method: "GET", Path: "/todos/1"},
		{Method: "GET", Path: "/todos/2"},
		{Method: "GET", Path: "/todos/3"},
		{Method: "GET", Path: "/todos/1"},
		{Method: "GET", Path: "/todos/2"},
		{Method: "GET", Path: "/todos/3"},
		{Method: "GET", Path: "/todos/1"},
		{Method: "GET", Path: "/todos/2"},
		{Method: "GET", Path: "/todos/3"},
		{Method: "GET", Path: "/todos/1"},
		{Method: "GET", Path: "/todos/2"},
		{Method: "GET", Path: "/todos/3"},
		{Method: "GET", Path: "/todos/1"},
		{Method: "GET", Path: "/todos/2"},
		{Method: "GET", Path: "/todos/3"},
	}

	responses, err := client.ParallelRequests(ctx, requests)
	if err != nil {
		t.Errorf("Parallel Request Error: %v", err.Error())
		return
	}

	for i, response := range responses {
		if err != nil {
			t.Errorf("Single Request Error: %v", err.Error())
			return
		}
		defer response.Body.Close()

		// Read raw body for additional processing or error messages
		body, err := io.ReadAll(response.Body)
		if err != nil {
			t.Errorf("Parse Body Error: %v", err.Error())
			return
		}
		t.Logf("Response %d: %+v\n", i+1, string(body))
	}
}
