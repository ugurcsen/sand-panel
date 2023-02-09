package users

import (
	"context"
	"github.com/ugurcsen/sand-panel/internal/api/grpc/users/pb"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
	"google.golang.org/grpc"
)

var _ pb.UserServiceServer = &Server{}

type Server struct {
	pb.UnimplementedUserServiceServer
	service ports.UserService
}

func (u *Server) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *Server) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *Server) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *Server) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (u *Server) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
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

	return &pb.GetResponse{Users: usersProto}, nil
}

func newServer(service ports.UserService, server grpc.Server) pb.UserServiceServer {
	return &Server{service: service}
}

func RegisterServer(service ports.UserService, server *grpc.Server) {
	pb.RegisterUserServiceServer(server, newServer(service, *server))
}
