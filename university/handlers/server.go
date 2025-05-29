package handlers

import (
	"university/cmd/proto"
	"university/service"
)

type UniversityHandler struct {
	proto.UnimplementedUniversityServiceServer
	Service service.UniversityService
}
