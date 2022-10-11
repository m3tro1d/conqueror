package transport

import (
	"context"

	conquerorapi "conqueror/api"
	"conqueror/pkg/conqueror/infrastructure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PublicGRPCServer interface {
	conquerorapi.ConquerorServer
}

func NewPublicGRPCServer(dependencyContainer infrastructure.DependencyContainer) PublicGRPCServer {
	return &publicGRPCServer{
		dependencyContainer: dependencyContainer,
	}
}

type publicGRPCServer struct {
	dependencyContainer infrastructure.DependencyContainer
}

func (s *publicGRPCServer) RegisterUser(ctx context.Context, request *conquerorapi.RegisterUserRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
