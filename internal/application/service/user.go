package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
)

// GetUsersByFilter returns the users that match the given filter.
func (a *Service) GetUsersByFilter(ctx context.Context, userFilter *pb.UserFilter) (*pb.Users, error) {
	return a.ServicePort.GetUsersByFilter(ctx, userFilter)
}

// CreateUser sends the given user to the service of the infrastructure layer for creating a new user.
func (a *Service) CreateUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	return a.ServicePort.CreateUser(ctx, user)
}

// UpdateUserRole sends the given user to the service of the infrastructure layer for updating user role.
func (a *Service) UpdateUserRole(ctx context.Context, user *pb.User) (*pb.User, error) {
	return a.ServicePort.UpdateUserRole(ctx, user)
}

// UpdateUserBase sends the given user to the service of the infrastructure layer for updating user base values.
func (a *Service) UpdateUserBase(ctx context.Context, user *pb.User) (*pb.User, error) {
	return a.ServicePort.UpdateUserBase(ctx, user)
}

// UpdateUserStatus sends the given user to the service of the infrastructure layer for updating user status.
func (a *Service) UpdateUserStatus(ctx context.Context, user *pb.User) (*pb.User, error) {
	return a.ServicePort.UpdateUserStatus(ctx, user)
}

// DeleteUser sends the given user to the service of the infrastructure layer for deleting data.
func (a *Service) DeleteUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	return a.ServicePort.DeleteUser(ctx, user)
}

// ChangePassword sends the given user password to the service of the infrastructure layer for changing user password.
func (a *Service) ChangePassword(ctx context.Context, userPassword *pb.UserPassword) error {
	return a.ServicePort.ChangePassword(ctx, userPassword)
}
