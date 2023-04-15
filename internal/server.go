package server

import (
	"context"
	"fmt"
	service "mihailbuslaev/sntt/pkg/api"
	"net/http"

	"golang.org/x/net/html"
)

type grpcServerImplement struct {
	service.UnimplementedRouteGuideServer
}

func (s *grpcServerImplement) GetCompany(ctx context.Context, req *service.GetCompanyRequest) (*service.GetCompanyResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.rusprofile.ru/search?query=%d", req.Inn))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 400 {
		return nil, fmt.Errorf("request to rusprofile failed, status is '%s'", resp.Status)
	}
	tkn := html.NewTokenizer(resp.Body)
	var data string
	for {
		tt := tkn.Next()
		var parseNext bool
		switch {
		case tt == html.ErrorToken:
			return nil, fmt.Errorf("error while parsing rusprofile")
		case tt == html.TextToken:
			t := tkn.Token()
			if parseNext {
				data = t.Data
				break
			}
			if t.Data == "Краткая справка" {
				parseNext = true
			}
		}
	}
	fmt.Println(data)
	return &service.GetCompanyResponse{}, nil
}

func NewGrpcServerImplement() *grpcServerImplement {
	return &grpcServerImplement{}
}
