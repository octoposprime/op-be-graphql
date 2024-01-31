package presentation

import (
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserFilterDto is a struct that provides filtering requested user data.
type UserFilterDto struct {
	ModelInputData *UserFilterInput
	PbData         *pb.UserFilter
}

// NewUserFilterDto creates a new *UserFilterDto.
func NewUserFilterDto(modelInputData *UserFilterInput) *UserFilterDto {
	return &UserFilterDto{
		ModelInputData: modelInputData,
		PbData:         new(pb.UserFilter),
	}
}

// ToPb converts *ModelInputData to *PbData and returns it.
func (u *UserFilterDto) ToPb() *pb.UserFilter {
	u.PbData.Id = u.ModelInputData.ID
	u.PbData.Username = u.ModelInputData.UserName
	u.PbData.Email = u.ModelInputData.Email
	if u.ModelInputData.UserType != nil {
		enumVal := pb.UserType(UserType_value[*u.ModelInputData.UserType])
		u.PbData.UserType = &enumVal
	}
	if u.ModelInputData.UserStatus != nil {
		enumVal := pb.UserStatus(UserStatus_value[*u.ModelInputData.UserStatus])
		u.PbData.UserStatus = &enumVal
	}
	u.PbData.Tags = u.ModelInputData.Tags
	u.PbData.FirstName = u.ModelInputData.FirstName
	u.PbData.LastName = u.ModelInputData.LastName
	if u.ModelInputData.CreatedAtFrom != nil && u.ModelInputData.CreatedAtTo != nil {
		u.PbData.CreatedAtFrom = timestamppb.New(*u.ModelInputData.CreatedAtFrom)
		u.PbData.CreatedAtTo = timestamppb.New(*u.ModelInputData.CreatedAtTo)
	}
	if u.ModelInputData.UpdatedAtFrom != nil && u.ModelInputData.UpdatedAtTo != nil {
		u.PbData.UpdatedAtFrom = timestamppb.New(*u.ModelInputData.UpdatedAtFrom)
		u.PbData.UpdatedAtTo = timestamppb.New(*u.ModelInputData.UpdatedAtTo)
	}
	u.PbData.SearchText = u.ModelInputData.SearchText
	u.PbData.SortType = u.ModelInputData.SortType
	if u.ModelInputData.SortField != nil {
		enumVal := pb.UserSortField(UserSortField_value[*u.ModelInputData.SortField])
		u.PbData.SortField = &enumVal
	}
	u.ModelInputData.Pagination = u.ModelInputData.Pagination.Validate()
	u.PbData.Limit = u.ModelInputData.Pagination.Limit
	u.PbData.Offset = u.ModelInputData.Pagination.Offset
	return u.PbData
}
