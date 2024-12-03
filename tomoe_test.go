package tomoe

import (
	"context"
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

	data, err := client.Do(ctx, opts)
	if err != nil {
		t.Errorf("Single Request Error: %v", err.Error())
		return
	}

	t.Logf("Success Response: %v", string(*data))
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

	data, err := client.ParallelRequests(ctx, requests)
	if err != nil {
		t.Errorf("Parallel Request Error: %v", err.Error())
		return
	}

	for i, result := range data {
		t.Logf("Response %d: %+v\n", i+1, string(*result))
	}
}
