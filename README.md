<!-- ![TOMOE](./tomoe.png) -->

# Tomoe: A Flexible and Fast HTTP Client Library for Go

**Tomoe** is a lightweight and extensible HTTP client library for Go, designed to simplify making HTTP requests while providing flexibility, speed, and robust error handling. It supports generics for type-safe responses, retries, parallel execution, and dynamic body data formats such as JSON, form-data, and file uploads.

---

## Features

- **Type-Safe Responses**: Use generics to ensure type safety and reduce boilerplate.
- **Retry Logic**: Automatic retry with exponential backoff for resilient requests.
- **Parallel Execution**: Perform concurrent HTTP requests efficiently using Goroutines.
- **Dynamic Body Support**: Seamlessly handle JSON, form-data, and file uploads.
- **Custom Headers and Query Parameters**: Easily configure requests for any use case.
- **Timeouts**: Control request execution time for better resource management.

---
## Installation
```bash
go get github.com/zakirkun/tomoe
```

## Usage
1. Basic Setup
```go 
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/zakirkun/tomoe"
)

func main() {
    client := NewClient("https://jsonplaceholder.typicode.com", 30*time.Second, 3, 5*time.Second, nil)
	ctx := context.Background()

	// Single request with retries
	opts := RequestOptions{
		Method: "GET",
		Path:   "/todos/1",
	}

	data, err := client.Do(ctx, opts)    
    if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

    
}
```


## Contributions
Contributions are welcome! Please fork the repository and submit a pull request.

## License
Tomoe is licensed under the MIT License. See the LICENSE file for details.

## Author
Created by Zakirkun. Inspired by a passion for elegant and efficient code.