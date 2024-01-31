package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
)

// UserServicePort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the other servies.
type UserServicePort interface {
	// GetUsersByFilter returns the users that match the given filter.
	GetUsersByFilter(ctx context.Context, userFilter *pb.UserFilter) (*pb.Users, error)

	// CreateUser sends the given user to the service of the infrastructure layer for creatimng user.
	CreateUser(ctx context.Context, user *pb.User) (*pb.User, error)

	// UpdateUserRole sends the given user to the service of the infrastructure layer for updating user role.
	UpdateUserRole(ctx context.Context, user *pb.User) (*pb.User, error)

	// UpdateUserBase sends the given user to the service of the infrastructure layer for updating user base values.
	UpdateUserBase(ctx context.Context, user *pb.User) (*pb.User, error)

	// UpdateUserStatus sends the given user to the service of the infrastructure layer for updating user status.
	UpdateUserStatus(ctx context.Context, user *pb.User) (*pb.User, error)

	// DeleteUser sends the given user to the service of the infrastructure layer for deleting user.
	DeleteUser(ctx context.Context, user *pb.User) (*pb.User, error)

	// ChangePassword sends the given user to the service of the infrastructure layer for updating user password.
	ChangePassword(ctx context.Context, userPassword *pb.UserPassword) error
}
