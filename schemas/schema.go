package schemas

import (
	"log"

	"github.com/graphql-go/graphql"
)

type Book struct {
	ID	int `json:"id"`
	Name string `json:"name"`
	Genre string `json:"genre"`
}

var bookList = []Book {
	Book {
		ID: 1,
		Name: "Concurrency in Go",
		Genre: "Programming",
	},
	Book {
		ID: 2,
		Name: "Programming Go",
		Genre: "Programming",
	},
}
/*
type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
} */

func createQueryType(bookType *graphql.Object) graphql.ObjectConfig {
	return graphql.ObjectConfig{Name: "RootQueryType", Fields: graphql.Fields{
		"book": &graphql.Field{
			Type: bookType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"]
				v, _ := id.(int)
				log.Printf("fetching post with id: %d", v)
				return fetchBookById(v), nil
			},
		},
	}}
}

func fetchBookById(id int) Book {
	for i := range bookList {
		if bookList[i].ID == id{
			return bookList[i]
		}
	}
	// Not found, return an empty object
	// TODO: What is the elegant approach - error object?
	return Book{
		ID:    0,
		Name:  "",
		Genre: "",
	}
}

func createBookType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"genre": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func ProcessQuery(query string) graphql.Result {

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(createQueryType(createBookType()))}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	params := graphql.Params{Schema: schema, RequestString: query}
	grp := graphql.Do(params)
	
	return *grp
}
