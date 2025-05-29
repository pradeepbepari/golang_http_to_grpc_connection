package handler

import (
	"context"
	"grpc-portal/cmd/proto"
	"grpc-portal/handler/protomodal"
)

func (s *Server) RegisterUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	userModalRequest := protomodal.ConvertUserProtoToUserModal(req)
	userData, err := s.UserService.UserRegister(ctx, *userModalRequest)
	if err != nil {
		return nil, err
	}
	userResponse := protomodal.ConvertUsermodalToUserProtoResponse(*userData)
	return userResponse, nil
}
