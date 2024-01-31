package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
)

// GetUsersByFilter returns the users that match the given filter.
func (a QueryAdapter) GetUsersByFilter(ctx context.Context, userFilter *pb.UserFilter) (*pb.Users, error) {
	return a.Service.GetUsersByFilter(ctx, userFilter)
}
