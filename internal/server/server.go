package server

import (
	"fmt"
	"net"
	"sync"

	"github.com/nanchano/hathor/internal/core"
	pb "github.com/nanchano/hathor/pb"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var once = &sync.Once{}

// server will implement the proto server interface to handle GRPC operations.
type server struct {
	logger  *slog.Logger
	service core.HathorService
	pb.UnimplementedHathorServer
}

// New creates a server given a logger and HathorService.
func New(l *slog.Logger, s core.HathorService) *server {
	return &server{
		logger:  l,
		service: s,
	}
}

// Start opens the GRPC server on the given host/port.
func (s *server) Start(host, port string) {
	once.Do(func() {
		url := fmt.Sprintf("%s:%s", host, port)
		s.logger.Info(fmt.Sprintf("Starting GRPC server on %s", url))
		grpcServer := grpc.NewServer()
		pb.RegisterHathorServer(grpcServer, s)
		reflection.Register(grpcServer)

		l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			s.logger.Error("Can't listen to the server")
			panic(err)
		}

		go func() {
			if err = grpcServer.Serve(l); err != nil {
				s.logger.Error("Failed to serve", "error", err)
			}
		}()

		s.logger.Info("GRPC server is live.")
	})
}
