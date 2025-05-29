package handlers

import (
	"university/cmd/proto"
	"university/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *UniversityHandler) MapUniversityToProto(university *models.University) *proto.UniversityResponse {
	return &proto.UniversityResponse{
		Id:              university.ID.Hex(),
		Name:            university.Name,
		Email:           university.Email,
		Contact:         university.Contact,
		Password:        university.Password,
		Logo:            university.Logo,
		Location:        university.Location,
		Website:         university.Website,
		City:            university.City,
		EstablishedYear: university.EstablishedYear,
		Colleges:        university.Colleges,
		Ranking:         university.Ranking,
		Programes:       university.Programs,
		CreatedAt:       timestamppb.New(university.CreatedAt),
		UpdatedAt:       timestamppb.New(university.UpdatedAt),
	}
}
func (u *UniversityHandler) MapProtoToUniversity(university *proto.UniversityRequest) models.University {
	return models.University{
		Name:            university.Name,
		Email:           university.Email,
		Contact:         university.Contact,
		Password:        university.Password,
		Logo:            university.Logo,
		Location:        university.Location,
		City:            university.City,
		Website:         university.Website,
		EstablishedYear: university.EstablishedYear,
		Colleges:        university.Colleges,
		Ranking:         university.Ranking,
		Programs:        university.Programs,
	}
}
func (u *UniversityHandler) MapProtoToUniversityByPagination(req *proto.ListUniversityRequest) models.Pagination {
	return models.Pagination{
		PageSize:  req.Pagination.Page,
		PageLimit: req.Pagination.Limit,
	}
}
func (u *UniversityHandler) MapUniversityToProtoList(universities []models.University) []*proto.UniversityResponse {
	universityList := make([]*proto.UniversityResponse, len(universities))
	for _, university := range universities {
		universityList = append(universityList, u.MapUniversityToProto(&university))
	}
	return universityList
}
func (u *UniversityHandler) MapProtoToUniversityById(req *proto.UniversityByIdRequest) models.University {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return models.University{}
	}
	return models.University{
		ID: id,
	}
}
func (u *UniversityHandler) MapProtoToUniversityByEmail(req *proto.UniversityByEmailRequest) models.University {
	return models.University{
		Email: req.Email,
	}
}

func (u *UniversityHandler) MapProtoToUniversityByName(req *proto.UniversityByNameRequest) models.University {
	return models.University{
		Name: req.Name,
	}
}
