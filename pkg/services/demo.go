package services

import (
	"context"
	"github.com/shiyongxi/go-gin-web/model/request"
	"github.com/shiyongxi/go-gin-web/model/response"
)

type (
	DemoService struct {
		context context.Context
	}
)

func NewDemoService(ctx context.Context) *DemoService {
	return &DemoService{
		context: ctx,
	}
}

func (svc *DemoService) GetInfo(formData *request.DemoRequest) (*response.DemoResponse, error) {
	return &response.DemoResponse{Name: formData.Name}, nil
}
