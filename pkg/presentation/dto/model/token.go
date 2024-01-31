package presentation

import pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/authentication"

// TokenDto is a struct that is used for authentication.
type TokenDto struct {
	ModelData      *Token
	ModelInputData *TokenInput
	PbData         *pb.Token
}

// NewTokenDto creates a new *TokenDto.
func NewTokenDto(modelInputData *TokenInput) *TokenDto {
	return &TokenDto{
		ModelData:      new(Token),
		ModelInputData: modelInputData,
		PbData:         new(pb.Token),
	}
}

// ToPb converts *ModelInputData to *PbData and returns it.
func (u *TokenDto) ToPb() *pb.Token {
	u.PbData.AuthenticationToken = u.ModelInputData.AuthenticationToken
	u.PbData.RefreshToken = u.ModelInputData.RefreshToken
	return u.PbData
}

// ToModel converts *PbData to *ModelData and returns it.
func (u *TokenDto) ToModel() *Token {
	u.ModelData.AuthenticationToken = u.PbData.AuthenticationToken
	u.ModelData.RefreshToken = u.PbData.RefreshToken
	return u.ModelData
}
