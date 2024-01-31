package presentation

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"

	mo "github.com/octoposprime/op-be-graphql/internal/domain/model/object"
	presentation "github.com/octoposprime/op-be-graphql/pkg/presentation/dto/model"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_user "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, user presentation.UserInput) (*presentation.User, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		user.ID = &userId
	}
	dtoData := presentation.NewUserDto(&user)
	resultData, err := r.CommandHandler.CreateUser(ctx, dtoData.ToPb())
	if err != nil {
		return nil, err
	}
	dtoData.PbData = resultData
	return dtoData.ToModel(), nil
}

// UpdateUserRole is the resolver for the updateUserRole field.
func (r *mutationResolver) UpdateUserRole(ctx context.Context, user presentation.UserInput) (*presentation.User, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		user.ID = &userId
	}
	dtoData := presentation.NewUserDto(&user)
	resultData, err := r.CommandHandler.UpdateUserRole(ctx, dtoData.ToPb())
	if err != nil {
		return nil, err
	}
	dtoData.PbData = resultData
	return dtoData.ToModel(), nil
}

// UpdateUserBase is the resolver for the updateUserBase field.
func (r *mutationResolver) UpdateUserBase(ctx context.Context, user presentation.UserInput) (*presentation.User, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		user.ID = &userId
	}
	dtoData := presentation.NewUserDto(&user)
	resultData, err := r.CommandHandler.UpdateUserBase(ctx, dtoData.ToPb())
	if err != nil {
		return nil, err
	}
	dtoData.PbData = resultData
	return dtoData.ToModel(), nil
}

// UpdateUserStatus is the resolver for the updateUserStatus field.
func (r *mutationResolver) UpdateUserStatus(ctx context.Context, user presentation.UserInput) (*presentation.User, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		user.ID = &userId
	}
	dtoData := presentation.NewUserDto(&user)
	resultData, err := r.CommandHandler.UpdateUserStatus(ctx, dtoData.ToPb())
	if err != nil {
		return nil, err
	}
	dtoData.PbData = resultData
	return dtoData.ToModel(), nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*presentation.User, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		id = userId
	}
	inData := new(presentation.UserInput)
	inData.ID = &id
	dtoData := presentation.NewUserDto(inData)
	resultData, err := r.CommandHandler.DeleteUser(ctx, dtoData.ToPb())
	if err != nil {
		return nil, err
	}
	dtoData.PbData = resultData
	return dtoData.ToModel(), nil
}

// ChangeUserPassword is the resolver for the changeUserPassword field.
func (r *mutationResolver) ChangeUserPassword(ctx context.Context, userPassword presentation.UserPasswordInput) (*bool, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		userPassword.UserID = userId
	}
	dtoData := presentation.NewUserPasswordDto(&userPassword)
	err := r.CommandHandler.ChangePassword(ctx, dtoData.ToPb())
	if err != nil {
		result := bool(false)
		return &result, err
	}
	result := bool(true)
	return &result, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*presentation.User, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		id = userId
	}
	var filter presentation.UserFilterInput
	filter.ID = &id
	dtoFilter := presentation.NewUserFilterDto(&filter)
	dtoData := presentation.NewUserDto(new(presentation.UserInput))
	resultDatas, err := r.QueryHandler.GetUsersByFilter(ctx, dtoFilter.ToPb())
	if err != nil {
		return nil, err
	}
	if len(resultDatas.Users) > 0 {
		dtoData.PbData = resultDatas.Users[0]
		return dtoData.ToModel(), nil
	} else {
		return nil, mo.ErrorUserNotFound
	}
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, filter *presentation.UserFilterInput) (*presentation.Users, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		filter.ID = &userId
	}
	var results presentation.Users
	if filter == nil {
		filter = &presentation.UserFilterInput{}
	}
	dtoFilter := presentation.NewUserFilterDto(filter)

	resultDatas, err := r.QueryHandler.GetUsersByFilter(ctx, dtoFilter.ToPb())
	if err != nil {
		return nil, err
	}

	for _, resultData := range resultDatas.Users {
		dtoData := presentation.NewUserDto(new(presentation.UserInput))
		dtoData.PbData = resultData
		results.Users = append(results.Users, dtoData.ToModel())
	}
	results.Total = int32(resultDatas.TotalRows)

	return &results, nil
}