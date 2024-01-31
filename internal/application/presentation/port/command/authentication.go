package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/authentication"
)

// AuthenticationCommandPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type AuthenticationCommandPort interface {
	// Login generates an authentication token if the given users are valid.
	Login(ctx context.Context, loginRequest *pb.LoginRequest) (*pb.Token, error)

	// Refresh regenerate an authentication token.
	Refresh(ctx context.Context, token *pb.Token) (*pb.Token, error)

	// Logout clears some footprints for the user.
	Logout(ctx context.Context, token *pb.Token) error
}
