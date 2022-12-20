package users

import (
	"context"
	"github.com/ugurcsen/sand-panel/internal/api/grpc/users/proto"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
	"google.golang.org/grpc"
)

type UsersServer struct {
	proto.UnimplementedUsersServer
	service ports.UserService
}

func (u *UsersServer) GetUsers(ctx context.Context, request *proto.GetUsersRequest) (*proto.GetUsersResponse, error) {
	users, err := u.service.List()
	if err != nil {
		return nil, err
	}

	usersProto := make([]*proto.User, len(users))
	for i, user := range users {
		if user == nil {
			continue
		}
		usersProto[i] = &proto.User{
			Name:  user.UserName,
			Email: *user.Email,
		}
	}

	return &proto.GetUsersResponse{Users: usersProto}, nil
}

func newUsersServer(service ports.UserService, server grpc.Server) proto.UsersServer {
	return &UsersServer{service: service}
}

func RegisterUsersServer(service ports.UserService, server *grpc.Server) {
	proto.RegisterUsersServer(server, newUsersServer(service, *server))
}
