# grpc-go

# Protoc

protoc --go_out=plugins=grpc:. proto/category/category.proto

# Demo Data
`
var categories []Category
	for i := 0; i < count; i++ {
		categories = append(categories, Category{
			ID:          int64(i),
			Title:       "This is Title " + strconv.Itoa(i),
			Description: "This is Description  " + strconv.Itoa(i),
		})
	}
	return categories`