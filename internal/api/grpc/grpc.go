package grpc

import (
	"fmt"
	"github.com/ugurcsen/sand-panel/internal/api/grpc/users"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
	GRPC "google.golang.org/grpc"
	"net"
	"reflect"
)

type grpc struct {
	server   *GRPC.Server
	services *Services
	options  *Options
}

type Services struct {
	UserService ports.UserService
}

type Options struct {
	Listener net.Listener
}

func (g *grpc) RegisterServices(services interface{}) error {
	if s, ok := services.(*Services); ok {
		g.services = s
		users.RegisterUsersServer(s.UserService, g.server)
		return nil
	}

	return fmt.Errorf("RegisterServices: only support %v", reflect.TypeOf(&Services{}))
}

func (g *grpc) Listen() error {
	return g.server.Serve(g.options.Listener)
}

func NewGrpc(options *Options) (ports.API, error) {
	return &grpc{
		server:  GRPC.NewServer(GRPC.EmptyServerOption{}),
		options: options,
	}, nil
}