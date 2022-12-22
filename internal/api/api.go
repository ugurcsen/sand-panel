package api

import (
	"fmt"
	"github.com/ugurcsen/sand-panel/internal/api/grpc"
	"github.com/ugurcsen/sand-panel/internal/api/rest"
	"github.com/ugurcsen/sand-panel/internal/core/services/user_service"
	"github.com/ugurcsen/sand-panel/internal/repositories/postgresql"
	"net"
)

func RunRest(port uint16) error {
	rps, _ := postgresql.NewUserRepository()
	srv := user_service.NewUserService(rps)
	tcp, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	defer tcp.Close()
	restApi, err := rest.NewRest(&rest.Options{
		Listener: tcp,
	})
	if err != nil {
		return err
	}

	err = restApi.RegisterServices(&rest.Services{
		UserService: srv,
	})
	if err != nil {
		return err
	}

	err = restApi.Listen()
	if err != nil {
		return err
	}

	return nil
}

func RunGrpc(port uint16) error {
	rps, _ := postgresql.NewUserRepository()
	srv := user_service.NewUserService(rps)
	tcp, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	defer tcp.Close()
	if err != nil {
		return err
	}
	grpcApi, err := grpc.NewGrpc(&grpc.Options{Listener: tcp})
	if err != nil {
		return err
	}

	err = grpcApi.RegisterServices(&grpc.Services{
		UserService: srv,
	})
	if err != nil {
		return err
	}

	err = grpcApi.Listen()
	if err != nil {
		return err
	}

	return nil
}
