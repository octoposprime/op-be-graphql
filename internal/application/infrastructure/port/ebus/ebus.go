package application

import (
	"context"

	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// EBusPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the event bus.
type EBusPort interface {
	// SetLogger sets logging call-back function
	SetLogger(LogFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error))

	//TODO: #20 Implement graphql notification ebus
	// Listen listens to the event bus and calls the given callBack function for each received any.
	Listen(ctx context.Context, channelName string, callBack func(channelName string, data any))
}
