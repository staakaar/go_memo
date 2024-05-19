//go:build ignore

package server

import (
	server "command-line-arguments/Users/iwamototakayuki/go-pro/go_memo/server/builder.go"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	svr := server.NewBuilder("localhost", 8888).Timeout(time.Minute).Logger(logger).Build()
	// builder := server.NewBuilder("localhost", 8888)
	// configureServer(builder)
	// svr := builder.Build()

	// svr := server.New(host, port, server.Config{
	// 	Timeout: time.Minute,
	// 	Logger:  nil,
	// })
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}

type Server struct {
	param serverParam
}

type serverParam struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func NewBuilder(host string, port int) *serverParam {
	return &serverParam{host: host, port: port}
}

func (sb *serverParam) Timeout(timeout time.Duration) *serverParam {
	sb.timeout = timeout
	return sb
}

func (sb *serverParam) Logger(logger *log.Logger) *serverParam {
	sb.logger = logger
	return sb
}

func (sb *serverParam) Build() *Server {
	svr := &Server{
		param: *sb,
	}
	return svr
}

func (s *Server) Start() error {
	if s.param.logger != nil {
		s.param.logger.Println("server started")
	}
	return nil
}
