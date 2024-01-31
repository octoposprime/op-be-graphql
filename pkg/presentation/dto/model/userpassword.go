package presentation

import (
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
)

// UserPasswordDto is a struct that represents the dto of a user password values.
type UserPasswordDto struct {
	ModelData      *UserPassword
	ModelInputData *UserPasswordInput
	PbData         *pb.UserPassword
}

// NewUserPasswordDto creates a new *UserPasswordDto.
func NewUserPasswordDto(modelInputData *UserPasswordInput) *UserPasswordDto {
	pbData := new(pb.UserPassword)
	return &UserPasswordDto{
		ModelData:      new(UserPassword),
		ModelInputData: modelInputData,
		PbData:         pbData,
	}
}

// ToPb converts *ModelInputData to *PbData and returns it.
func (u *UserPasswordDto) ToPb() *pb.UserPassword {
	if u.ModelInputData.ID != nil {
		u.PbData.Id = *u.ModelInputData.ID
	}
	u.PbData.UserId = u.ModelInputData.UserID
	u.PbData.Password = u.ModelInputData.Password
	return u.PbData
}

// ToModel converts *PbData to *ModelData and returns it.
func (u *UserPasswordDto) ToModel() *UserPassword {
	u.ModelData.ID = u.PbData.Id
	u.ModelData.UserID = u.PbData.UserId
	u.ModelData.Password = u.PbData.Password
	u.ModelData.PasswordStatus = PasswordStatus(PasswordStatus_name[int32(u.PbData.PasswordStatus)])
	return u.ModelData
}
