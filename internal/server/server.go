package server

import (
	"context"
	service "mihailbuslaev/pb-wrapper/pkg/api"
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
