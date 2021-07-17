package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	ip       string
	port     string
	listener net.Listener
}

func New() (*Server, error) {
	defaultPort := "4000"
	listener, err := net.Listen("tcp", fmt.Sprintf(":"+defaultPort))
	if err != nil {
		return nil, fmt.Errorf("error creating listener: %v", err)
	}

	return &Server{
		ip:       listener.Addr().(*net.TCPAddr).IP.String(),
		port:     defaultPort,
		listener: listener,
	}, nil

}

func NewWithPort(port string) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":"+port))
	if err != nil {
		return nil, fmt.Errorf("error creating listener: %v", err)
	}

	return &Server{
		ip:       listener.Addr().(*net.TCPAddr).IP.String(),
		port:     port,
		listener: listener,
	}, nil
}

func (s *Server) ServeHTTP(ctx context.Context, server *http.Server) {

	ctx, cancel := context.WithCancel(ctx)

	// goroutine to listen for os interrupts
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		fmt.Println("signal caught. shutting server down...")
		cancel()
	}()

	done := make(chan struct{})
	go func() {
		<-ctx.Done()

		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("failed to shutdown: %v\n", err)
		}
		close(done)
	}()

	if err := server.Serve(s.listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("error serving: %v\n", err)
	}

	<-done
}
