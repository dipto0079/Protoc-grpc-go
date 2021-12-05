package category

import (
	"context"
	cpb "grpc-category/proto/category"

	"google.golang.org/grpc"
)

type Client struct{
	client cpb.CategoryServiceClient
}

func NewClient(conn grpc.ClientConnInterface) Client{
	return Client{
		client: cpb.NewCategoryServiceClient(conn),
	}
}

func (c *Client)GetCategory(id int64)(*cpb.GetCategoryResponse,error){
	
	return c.client.GetCategory(context.Background(),&cpb.GetCategoryRequest{
		ID: id,
	})
}