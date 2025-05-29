package handlers

import (
	"context"
	"university/cmd/proto"
)

func (u *UniversityHandler) CreateUniversity(ctx context.Context, req *proto.UniversityRequest) (*proto.UniversityResponse, error) {
	request := u.MapProtoToUniversity(req)
	university, err := u.Service.CreateUniversity(ctx, request)
	if err != nil {
		return nil, err
	}
	response := u.MapUniversityToProto(university)
	return response, nil
}
func (u *UniversityHandler) ListUniversity(ctx context.Context, req *proto.ListUniversityRequest) (*proto.ListUniversityResponse, error) {
	request := u.MapProtoToUniversityByPagination(req)
	universities, count, err := u.Service.ListUniversity(ctx, request)
	if err != nil {
		return nil, err
	}
	response := &proto.ListUniversityResponse{
		University: u.MapUniversityToProtoList(universities),
		Total:      count,
		Pagination: &proto.Pagination{
			Page:  req.Pagination.Page,
			Limit: req.Pagination.Limit,
		}}

	return response, nil

}
func (u *UniversityHandler) UpdateUniversity(ctx context.Context, req *proto.UniversityRequest) (*proto.UniversityResponse, error) {
	request := u.MapProtoToUniversity(req)
	university, err := u.Service.UpdateUniversity(ctx, request)
	if err != nil {
		return nil, err
	}
	response := u.MapUniversityToProto(university)
	return response, nil

}

func (u *UniversityHandler) GetUniversityById(ctx context.Context, req *proto.UniversityByIdRequest) (*proto.UniversityResponse, error) {
	request := u.MapProtoToUniversityById(req)
	university, err := u.Service.GetUniversityById(ctx, request.ID)
	if err != nil {
		return nil, err
	}
	response := u.MapUniversityToProto(university)
	return response, nil

}
func (u *UniversityHandler) GetUniversityByEmail(ctx context.Context, req *proto.UniversityByEmailRequest) (*proto.UniversityResponse, error) {
	request := u.MapProtoToUniversityByEmail(req)
	university, err := u.Service.GetUniversityByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}
	response := u.MapUniversityToProto(university)
	return response, nil

}

func (u *UniversityHandler) GetUniversityByName(ctx context.Context, req *proto.UniversityByNameRequest) (*proto.UniversityResponse, error) {
	request := u.MapProtoToUniversityByName(req)
	university, err := u.Service.GetUniversityByName(ctx, request.Name)
	if err != nil {
		return nil, err
	}
	response := u.MapUniversityToProto(university)
	return response, nil

}
