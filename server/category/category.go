package category

import (
	"context"
	cpb "grpc-category/proto/category"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct{}

type Category struct {
	ID          int64
	Title       string
	Description string
}

var categorys = []Category{
	{
		ID:          1,
		Title:       "This is title 1",
		Description: "This is Description 1",
	},
	{
		ID:          2,
		Title:       "This is title 2",
		Description: "This is Description 2",
	},
	{
		ID:          3,
		Title:       "This is title 3",
		Description: "This is Description 3",
	},
}

func (s *Server) GetCategory(ctx context.Context, req *cpb.GetCategoryRequest) (*cpb.GetCategoryResponse, error) {
	
	log.Printf("Category ID:%d", req.GetID())

	var category Category

	for _, c := range categorys {
		if c.ID == req.GetID() {
			category = c
			break
		}
	}

	if category.ID == 0{
		return &cpb.GetCategoryResponse{},status.Errorf(codes.NotFound,"invalid id")
	}

	return &cpb.GetCategoryResponse{
		ID: category.ID,
		Title: category.Title,
		Description: category.Description,
	},nil

}
