package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/authentication"
)

// Login generates an authentication token if the given users are valid.
func (a *Service) Login(ctx context.Context, loginRequest *pb.LoginRequest) (*pb.Token, error) {
	return a.ServicePort.Login(ctx, loginRequest)
}

// Refresh regenerate an authentication token.
func (a *Service) Refresh(ctx context.Context, token *pb.Token) (*pb.Token, error) {
	return a.ServicePort.Refresh(ctx, token)
}

// Logout clears some footprints for the user.
func (a *Service) Logout(ctx context.Context, token *pb.Token) error {
	return a.ServicePort.Logout(ctx, token)
}
