package main

import (
	"grpc-category/client/category"
	"log"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to Connect:%s", err)
	}
	categoryClient := category.NewClient(conn)

	res, err := categoryClient.GetCategory(3)

	if err != nil {
		log.Fatalf("error while calling GetCategory:%s", err)
	}
	
	log.Printf("Response", res)

	if err := categoryClient.GetCategorys(); err != nil {
		log.Fatalf("error while calling GetCategorys:%s", err)
	}

}
