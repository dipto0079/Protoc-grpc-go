package category

import (
	"context"
	cpb "grpc-category/proto/category"
	"io"
	"log"

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

func (c *Client)GetCategorys()error{
	
	stream,err := c.client.GetCategorys(context.Background(),&cpb.GetCategorysRequest{})
	if err!= nil{
		return err
	}

	for{
		res,err:=stream.Recv()
		if err !=nil{
			if err == io.EOF{
				return nil
			}
			return err
		}
		log.Printf("Received ID: %d",res.GetID())
	}
}