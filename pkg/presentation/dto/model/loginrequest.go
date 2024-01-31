package presentation

import (
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/authentication"
)

// LoginRequestDto is a struct that represents the dto of a user basic values.
type LoginRequestDto struct {
	ModelData      *LoginRequest
	ModelInputData *LoginRequestInput
	PbData         *pb.LoginRequest
}

// NewLoginRequestDto creates a new *LoginRequestDto.
func NewLoginRequestDto(modelInputData *LoginRequestInput) *LoginRequestDto {
	pbData := new(pb.LoginRequest)
	return &LoginRequestDto{
		ModelData:      new(LoginRequest),
		ModelInputData: modelInputData,
		PbData:         pbData,
	}
}

// ToPb converts *ModelInputData to *PbData and returns it.
func (u *LoginRequestDto) ToPb() *pb.LoginRequest {
	if u.ModelInputData.UserName != nil {
		u.PbData.Username = *u.ModelInputData.UserName
	}
	if u.ModelInputData.Email != nil {
		u.PbData.Email = *u.ModelInputData.Email
	}
	if u.ModelInputData.Password != nil {
		u.PbData.Password = *u.ModelInputData.Password
	}
	return u.PbData
}

// ToModel converts *PbData to *ModelData and returns it.
func (u *LoginRequestDto) ToModel() *LoginRequest {
	u.ModelData.UserName = u.PbData.Username
	u.ModelData.Email = u.PbData.Email
	u.ModelData.Password = u.PbData.Password
	return u.ModelData
}
