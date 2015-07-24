# go-semaphore
A simple semaphore type

# Usage
```go
semaphore := semaphore.New(8)

// Start 100 goroutines
for i = 0; i < 100; i++ {
	go func doWork() {
		// Make a network request
		makeRequest()
	}()
}

func makeRequest() {
	// Limit maximum concurrent requests to 8 with the semaphore
	semaphore.Yield()
	defer semaphore.Unyield()

	// Make request...
}
```

# Notes
Any ideas on how to test something like this are welcome.