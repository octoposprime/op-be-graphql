package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
)

// UserCommandPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type UserCommandPort interface {
	// CreateUser sends the given user to the application layer for creating a new user.
	CreateUser(ctx context.Context, user *pb.User) (*pb.User, error)

	// UpdateUserRole sends the given user to the application layer for updating user role.
	UpdateUserRole(ctx context.Context, user *pb.User) (*pb.User, error)

	// UpdateUserBase sends the given user to the application layer for updating user base values.
	UpdateUserBase(ctx context.Context, user *pb.User) (*pb.User, error)

	// UpdateUserStatus sends the given user to the application layer for updating user status.
	UpdateUserStatus(ctx context.Context, user *pb.User) (*pb.User, error)

	// DeleteUser sends the given user to the application layer for deleting data.
	DeleteUser(ctx context.Context, user *pb.User) (*pb.User, error)

	// ChangePassword sends the given user to the application layer for changing user password.
	ChangePassword(ctx context.Context, userPassword *pb.UserPassword) error
}
