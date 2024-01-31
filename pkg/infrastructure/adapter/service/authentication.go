package infrastructure

import (
	"context"

	me "github.com/octoposprime/op-be-graphql/internal/domain/model/entity"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/authentication"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tconfig "github.com/octoposprime/op-be-shared/tool/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Login generates an authentication token if the given users are valid.
func (a ServiceAdapter) Login(ctx context.Context, loginRequest *pb.LoginRequest) (*pb.Token, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.UserHost+":"+tconfig.GetInternalConfigInstance().Grpc.UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Login", userId, err.Error()))
		return &pb.Token{}, err
	}

	pbResult, err := pb.NewAuthenticationSvcClient(conn).Login(ctx, loginRequest)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Login", userId, err.Error()))
		return &pb.Token{}, err
	}
	return pbResult, nil
}

// Refresh regenerate an authentication token.
func (a ServiceAdapter) Refresh(ctx context.Context, token *pb.Token) (*pb.Token, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.UserHost+":"+tconfig.GetInternalConfigInstance().Grpc.UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Refresh", userId, err.Error()))
		return &pb.Token{}, err
	}

	pbResult, err := pb.NewAuthenticationSvcClient(conn).Refresh(ctx, token)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Refresh", userId, err.Error()))
		return &pb.Token{}, err
	}
	return pbResult, nil
}

// Logout clears some footprints for the user.
func (a ServiceAdapter) Logout(ctx context.Context, token *pb.Token) error {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.UserHost+":"+tconfig.GetInternalConfigInstance().Grpc.UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Logout", userId, err.Error()))
		return err
	}

	_, err = pb.NewAuthenticationSvcClient(conn).Logout(ctx, token)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Logout", userId, err.Error()))
		return err
	}
	return nil
}
