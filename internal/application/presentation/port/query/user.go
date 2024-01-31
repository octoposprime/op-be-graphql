package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
)

// QueryPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type UserQueryPort interface {
	// GetUsersByFilter returns the users that match the given filter.
	GetUsersByFilter(ctx context.Context, userFilter *pb.UserFilter) (*pb.Users, error)
}
