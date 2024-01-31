package presentation

import (
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/error"
)

// ErrorDto is a struct that represents the dto of a built-in error.
type ErrorDto struct {
	ModelData *Error
	PbData    *pb.Error
}

// NewErrorDto creates a new *ErrorDto.
func NewErrorDto() *ErrorDto {
	pbData := new(pb.Error)
	return &ErrorDto{
		ModelData: new(Error),
		PbData:    pbData,
	}
}

// ToModel converts *PbData to *ModelData and returns it.
func (u *ErrorDto) ToModel() *Error {
	u.ModelData.Error = u.PbData.Error
	return u.ModelData
}
