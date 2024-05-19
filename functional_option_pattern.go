//go:build ignore

package main

import (
	"log"
	"os"
	"time"

	"github.com/mattn/awesome-server/server"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

type Option func(*Server)

func main() {
	f, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	svr := server.New("localhost", 8888, server.WithTiomeout(time.Minute), server.WithLogger(logger))
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}

func New(host string, port int, options ...Option) *Server {
	svr := &Server{
		host: host,
		port: port,
	}

	for _, opt := range options {
		opt(svr)
	}
	return svr
}

func (s *Server) Start() error {
	return nil
}

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}
func WithLogger(logger *log.Logger) func(*Server) {
	return func(s *Server) {
		s.logger = logger
	}
}
