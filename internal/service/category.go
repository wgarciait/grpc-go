package service

import (
	"github.com/wgarciait/grpc-go/internal/database"
	"github.com/wgarciait/grpc-go/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}

func (c *CategoryService) CreateCategory(context.Context, in *CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)

	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category {
		Id:				category.ID,
		Name:			category.Name,
		Description:	category.Description,
	}

	return &pb.CategoryResponse {
		Category: categoryResponse,
	}, nil
}

