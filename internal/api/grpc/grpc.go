package grpc

import (
	"fmt"
	"github.com/ugurcsen/sand-panel/internal/api/grpc/users"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
	GRPC "google.golang.org/grpc"
	"net"
	"reflect"
)

// grpc is the grpc api
type grpc struct {
	server   *GRPC.Server
	services *Services
	options  *Options
}

// Services is the services that the grpc api will use
type Services struct {
	UserService ports.UserService
}

// Options is the options for the grpc api
type Options struct {
	Listener net.Listener
}

// RegisterServices registers the services to the grpc api
func (g *grpc) RegisterServices(services interface{}) error {
	if s, ok := services.(*Services); ok {
		g.services = s
		users.RegisterUsersServer(s.UserService, g.server)
		return nil
	}

	return fmt.Errorf("RegisterServices: only support %v", reflect.TypeOf(&Services{}))
}

// Listen starts the grpc api
func (g *grpc) Listen() error {
	return g.server.Serve(g.options.Listener)
}

// NewGrpc creates a new grpc api
func NewGrpc(options *Options) (*grpc, error) {
	return &grpc{
		server:  GRPC.NewServer(GRPC.EmptyServerOption{}),
		options: options,
	}, nil
}
