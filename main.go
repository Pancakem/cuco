package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())

	server, err := NewWithPort("5000")
	if err != nil {
		fmt.Println("%w", err)
	}

	srv := http.Server {
		Handler: Handler(),
		ReadTimeout: 1 * time.Minute,
	}

	server.ServeHTTP(ctx, &srv)
}
