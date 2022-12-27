package users

import (
	"context"
	"github.com/ugurcsen/sand-panel/internal/api/grpc/users/pb"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	service ports.UserService
}

func (u *Server) GetUsers(ctx context.Context, request *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, err := u.service.List()
	if err != nil {
		return nil, err
	}

	usersProto := make([]*pb.User, len(users))
	for i, user := range users {
		if user == nil {
			continue
		}
		usersProto[i] = &pb.User{
			Name:  user.UserName,
			Email: *user.Email,
		}
	}

	return &pb.GetUsersResponse{Users: usersProto}, nil
}

func newServer(service ports.UserService, server grpc.Server) pb.UserServiceServer {
	return &Server{service: service}
}

func RegisterServer(service ports.UserService, server *grpc.Server) {
	pb.RegisterUserServiceServer(server, newServer(service, *server))
}
