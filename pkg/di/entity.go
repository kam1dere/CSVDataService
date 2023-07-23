package di

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	Ln   net.Listener
	Serv *grpc.Server
}

func (g *GrpcServer) Shutdown() error {
	g.Serv.Stop()

	err := g.Ln.Close()
	if err != nil {
		return fmt.Errorf("shutdown error: %w", err)
	}

	return nil
}

func (g *GrpcServer) Start() error {
	err := g.Serv.Serve(g.Ln)
	if err != nil {
		return fmt.Errorf("start error: %w", err)
	}

	return nil
}

type Shutdownable interface {
	Shutdown() error
}

type Closeable interface {
	Close() error
}
