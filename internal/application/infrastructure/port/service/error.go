package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/error"
)

// UserServicePort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the other servies.
type ErrorServicePort interface {
	// GetErrors returns the built-in errors.
	GetErrors(ctx context.Context) (*pb.Errors, error)
}
