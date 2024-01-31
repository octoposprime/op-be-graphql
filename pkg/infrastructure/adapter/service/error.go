package infrastructure

import (
	"context"

	me "github.com/octoposprime/op-be-graphql/internal/domain/model/entity"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/error"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tconfig "github.com/octoposprime/op-be-shared/tool/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// GetErrors returns the built-in errors.
func (a ServiceAdapter) GetErrors(ctx context.Context) (*pb.Errors, error) {
	var pbResults pb.Errors
	loggingErrors := a.GetErrorsDynamic(ctx, "GetLoggingErrors", tconfig.GetInternalConfigInstance().Grpc.LoggerHost, tconfig.GetInternalConfigInstance().Grpc.LoggerPort)
	userErrors := a.GetErrorsDynamic(ctx, "GetUserErrors", tconfig.GetInternalConfigInstance().Grpc.UserHost, tconfig.GetInternalConfigInstance().Grpc.UserPort)
	pbResults.Errors = append(pbResults.Errors, loggingErrors.Errors...)
	pbResults.Errors = append(pbResults.Errors, userErrors.Errors...)
	return &pbResults, nil
}

func (a ServiceAdapter) GetErrorsDynamic(ctx context.Context, msg string, host string, port string) *pb.Errors {
	conn, err := grpc.Dial(host+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, msg, userId, err.Error()))
		return &pb.Errors{}
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResults, err := pb.NewErorrSvcClient(conn).GetErrors(ctx, &pb.ErrorRequest{})
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, msg, userId, err.Error()))
		return &pb.Errors{}
	}
	return pbResults
}
