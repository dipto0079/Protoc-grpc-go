package category

import (
	"context"
	"log"
	"strconv"
	"time"

	cpb "grpc-category/proto/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct{}

type Category struct {
	ID          int64
	Title       string
	Description string
}



func (s *Server) GetCategory(ctx context.Context, req *cpb.GetCategoryRequest) (*cpb.GetCategoryResponse, error) {
	log.Printf("Category ID:%d", req.GetID())
	var category Category
	categorys := s.getDummyCategory(10)
	for _, c := range categorys {
		if c.ID == req.GetID() {
			category = c
			break
		}
	}
	if category.ID == 0 {
		return &cpb.GetCategoryResponse{}, status.Errorf(codes.NotFound, "invalid id")
	}
	return &cpb.GetCategoryResponse{
		ID:          category.ID,
		Title:       category.Title,
		Description: category.Description,
	}, nil
}

func (s *Server) getDummyCategory(count int) []Category {
	var categories []Category
	for i := 0; i < count; i++ {
		categories = append(categories, Category{
			ID:          int64(i),
			Title:       "This is Title " + strconv.Itoa(i),
			Description: "This is Description  " + strconv.Itoa(i),
		})
	}
	return categories
}

func (s *Server) GetCategorys(req *cpb.GetCategorysRequest, stream cpb.CategoryService_GetCategorysServer) error {
	categorys := s.getDummyCategory(5)
	for _, c := range categorys {
		err := stream.Send(&cpb.GetCategoryResponse{
			ID:          c.ID,
			Title:       c.Title,
			Description: c.Description,
		})
		if err != nil {
			return status.Error(codes.Internal, "failed to sand")
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}
