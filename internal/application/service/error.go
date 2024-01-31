package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/error"
)

// GetErrors returns the built-in errors.
func (a *Service) GetErrors(ctx context.Context) (*pb.Errors, error) {
	return a.ServicePort.GetErrors(ctx)
}
