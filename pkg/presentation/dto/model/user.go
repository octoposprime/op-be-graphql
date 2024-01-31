package presentation

import (
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
)

// UserDto is a struct that represents the dto of a user basic values.
type UserDto struct {
	ModelData      *User
	ModelInputData *UserInput
	PbData         *pb.User
}

// NewUserDto creates a new *UserDto.
func NewUserDto(modelInputData *UserInput) *UserDto {
	pbData := new(pb.User)
	return &UserDto{
		ModelData:      new(User),
		ModelInputData: modelInputData,
		PbData:         pbData,
	}
}

// ToPb converts *ModelInputData to *PbData and returns it.
func (u *UserDto) ToPb() *pb.User {
	if u.ModelInputData.ID != nil {
		u.PbData.Id = *u.ModelInputData.ID
	}
	if u.ModelInputData.UserName != nil {
		u.PbData.Username = *u.ModelInputData.UserName
	}
	if u.ModelInputData.Email != nil {
		u.PbData.Email = *u.ModelInputData.Email
	}
	if u.ModelInputData.Role != nil {
		u.PbData.Role = *u.ModelInputData.Role
	}
	if u.ModelInputData.UserType != nil {
		u.PbData.UserType = pb.UserType(UserType_value[*u.ModelInputData.UserType])
	}
	if u.ModelInputData.UserStatus != nil {
		u.PbData.UserStatus = pb.UserStatus(UserStatus_value[*u.ModelInputData.UserStatus])
	}
	if u.ModelInputData.UserBase != nil {
		if u.ModelInputData.UserBase.Tags != nil {
			u.PbData.Tags = u.ModelInputData.UserBase.Tags
		}
		if u.ModelInputData.UserBase.FirstName != nil {
			u.PbData.FirstName = *u.ModelInputData.UserBase.FirstName
		}
		if u.ModelInputData.UserBase.LastName != nil {
			u.PbData.LastName = *u.ModelInputData.UserBase.LastName
		}
	}
	return u.PbData
}

// ToModel converts *PbData to *ModelData and returns it.
func (u *UserDto) ToModel() *User {
	u.ModelData.ID = u.PbData.Id
	u.ModelData.UserName = u.PbData.Username
	u.ModelData.Email = u.PbData.Email
	u.ModelData.Role = u.PbData.Role
	u.ModelData.UserType = UserType(UserType_name[int32(u.PbData.UserType)])
	u.ModelData.UserStatus = UserStatus(UserStatus_name[int32(u.PbData.UserStatus)])
	u.ModelData.UserBase.Tags = u.PbData.Tags
	u.ModelData.UserBase.FirstName = u.PbData.FirstName
	u.ModelData.UserBase.LastName = u.PbData.LastName

	// Only for view
	u.ModelData.CreatedAt = u.PbData.CreatedAt.AsTime()
	u.ModelData.UpdatedAt = u.PbData.UpdatedAt.AsTime()
	return u.ModelData
}
