package protomodal

import (
	"grpc-portal/cmd/proto"
	"grpc-portal/model"
)

func ConvertUserProtoToUserModal(req *proto.UserRequest) *model.User {
	return &model.User{
		Name:     req.Name,
		Password: req.Password,
		Country:  req.Country,
		Email:    req.Email,
		Contact:  req.Contact,
		Address:  req.Address,
	}
}
func ConvertUsermodalToUserProtoResponse(req model.User) *proto.UserResponse {
	return &proto.UserResponse{
		Uuid: req.Id.String(),
	}
}
