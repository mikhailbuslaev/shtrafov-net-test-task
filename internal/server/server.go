package server

import (
	"context"
	service "mihailbuslaev/pb-wrapper/pkg/api/v1"
)

type grpcServerImplement struct {
	service.UnimplementedRouteGuideServer
}

func (s *grpcServerImplement) GetCompany(ctx context.Context, req *service.GetCompanyRequest) (*service.GetCompanyResponse, error) {
	return &service.GetCompanyResponse{}, nil
}

func NewGrpcServerImplement() *grpcServerImplement {
	return &grpcServerImplement{}
}
