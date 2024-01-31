package infrastructure

import (
	"context"

	me "github.com/octoposprime/op-be-graphql/internal/domain/model/entity"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
	tconfig "github.com/octoposprime/op-be-shared/tool/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// GetUsersByFilter returns the users that match the given filter.
func (a ServiceAdapter) GetUsersByFilter(ctx context.Context, userFilter *pb.UserFilter) (*pb.Users, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.UserHost+":"+tconfig.GetInternalConfigInstance().Grpc.UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetUsersByFilter", userId, err.Error()))
		return &pb.Users{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResults, err := pb.NewUserSvcClient(conn).GetUsersByFilter(ctx, userFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetUsersByFilter", userId, err.Error()))
		return &pb.Users{}, err
	}
	return pbResults, nil
}

// CreateUser insert a new user in the databse.
func (a ServiceAdapter) CreateUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.UserHost+":"+tconfig.GetInternalConfigInstance().Grpc.UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreateUser", userId, err.Error()))
		return &pb.User{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewUserSvcClient(conn).CreateUser(ctx, user)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreateUser", userId, err.Error()))
		return &pb.User{}, err
	}
	return pbResult, nil
}

// UpdateUserRole sends the given user to the application layer for updating user role.
func (a ServiceAdapter) UpdateUserRole(ctx context.Context, user *pb.User) (*pb.User, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.UserHost+":"+tconfig.GetInternalConfigInstance().Grpc.UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserRole", userId, err.Error()))
		return &pb.User{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewUserSvcClient(conn).UpdateUserRole(ctx, user)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserRole", userId, err.Error()))
		return &pb.User{}, err
	}
	return pbResult, nil
}

// UpdateUserBase sends the given user to the application layer for updating user base values.
func (a ServiceAdapter) UpdateUserBase(ctx context.Context, user *pb.User) (*pb.User, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.UserHost+":"+tconfig.GetInternalConfigInstance().Grpc.UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return &pb.User{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewUserSvcClient(conn).UpdateUserBase(ctx, user)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return &pb.User{}, err
	}
	return pbResult, nil
}

// UpdateUserStatus sends the given user to the application layer for updating user status.
func (a ServiceAdapter) UpdateUserStatus(ctx context.Context, user *pb.User) (*pb.User, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.UserHost+":"+tconfig.GetInternalConfigInstance().Grpc.UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return &pb.User{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewUserSvcClient(conn).UpdateUserStatus(ctx, user)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return &pb.User{}, err
	}
	return pbResult, nil
}

// DeleteUser soft-deletes the given user in the database.
func (a ServiceAdapter) DeleteUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.UserHost+":"+tconfig.GetInternalConfigInstance().Grpc.UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeleteUser", userId, err.Error()))
		return &pb.User{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewUserSvcClient(conn).DeleteUser(ctx, user)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeleteUser", userId, err.Error()))
		return &pb.User{}, err
	}
	return pbResult, nil
}

// ChangePassword changes the poassword of the given user in the database.
func (a ServiceAdapter) ChangePassword(ctx context.Context, userPassword *pb.UserPassword) error {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.UserHost+":"+tconfig.GetInternalConfigInstance().Grpc.UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "ChangePassword", userId, err.Error()))
		return err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	_, err = pb.NewUserSvcClient(conn).ChangePassword(ctx, userPassword)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "ChangePassword", userId, err.Error()))
		return err
	}
	return nil
}
