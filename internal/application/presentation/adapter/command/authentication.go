package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/authentication"
)

// Login generates an authentication token if the given users are valid.
func (a CommandAdapter) Login(ctx context.Context, loginRequest *pb.LoginRequest) (*pb.Token, error) {
	return a.Service.Login(ctx, loginRequest)
}

// Refresh regenerate an authentication token.
func (a CommandAdapter) Refresh(ctx context.Context, token *pb.Token) (*pb.Token, error) {
	return a.Service.Refresh(ctx, token)
}

// Logout clears some footprints for the user.
func (a CommandAdapter) Logout(ctx context.Context, token *pb.Token) error {
	return a.Service.Logout(ctx, token)
}
