package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
)

// CreateUser sends the given user to the application layer for creating a new user.
func (a CommandAdapter) CreateUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	return a.Service.CreateUser(ctx, user)
}

// UpdateUserRole sends the given user to the application layer for updating user role.
func (a CommandAdapter) UpdateUserRole(ctx context.Context, user *pb.User) (*pb.User, error) {
	return a.Service.UpdateUserRole(ctx, user)
}

// UpdateUserBase sends the given user to the application layer for updating user values.
func (a CommandAdapter) UpdateUserBase(ctx context.Context, user *pb.User) (*pb.User, error) {
	return a.Service.UpdateUserBase(ctx, user)
}

// UpdateUserStatus sends the given user to the application layer for updating user status.
func (a CommandAdapter) UpdateUserStatus(ctx context.Context, user *pb.User) (*pb.User, error) {
	return a.Service.UpdateUserStatus(ctx, user)
}

// DeleteUser sends the given user to the application layer for deleting data.
func (a CommandAdapter) DeleteUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	return a.Service.DeleteUser(ctx, user)
}

// ChangePassword sends the given user password to the application layer for changing user password.
func (a CommandAdapter) ChangePassword(ctx context.Context, userPassword *pb.UserPassword) error {
	return a.Service.ChangePassword(ctx, userPassword)
}
