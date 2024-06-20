package main

import (
    "encoding/json"
    "fmt"
    "github.com/graphql-go/graphql"
)

type Tutorial struct {
    ID int
    Title string
    Author Author
    Comments []Comment
}

type Author struct {
    Name string
    Tutorials []int
}

type Comment struct {
    Body string
}

func populate() []Tutorial {
    author := &Author{Name: "Elliot Forbes", Tutorials: []int{1}}
    tutorial := Tutorial{
        ID: 1,
        Title: "Go GraphQL Tutorial",
        Author: *author,
        Comments: []Comment{
            Comment{Body: "First Comment"},
        },
    }
    var tutorials []Tutorial
    tutorials = append(tutorials, tutorial)

    return tutorials
}

func main() {
    fmt.Println("graphql tutorial")

    var commentType = graphql.NewObject(
        graphql.ObjectConfig{
            Name: "Comment",
            Fields: graphql.Fields{
                "body": &graphql.Field{
                    Type: graphql.String,
                },
            },
        },
    )

    var authorType = graphql.NewObject(
        graphql.ObjectConfig{
            Name: "Author",
            Fields: graphql.Fields{
                "name": &graphql.Field{
                    Type: graphql.String,
                },
                "tutorials": &graphql.Field{
                    Type: graphql.NewList(graphql.Int),
                },
            },
        },
    )

    var tutorialType = graphql.NewObject(
        graphql.ObjectConfig{
            Name: "Tutorial",
            Fields: graphql.Fields{
                "id": &graphql.Field{
                    Type: graphql.Int,
                },
                "title": &graphql.Field{
                    Type: graphql.String,
                },
                "author": &graphql.Field{
                    Type: authorType,
                },
                "comments": &graphql.Field{
                    Type: graphql.NewList(commentType),
                },
            },
        },
    )

    // schema definition
    fields := graphql.Fields{
        "hello": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return "world", nil
            },
        },
    }

    rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
    schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
    schema, err := graphql.NewSchema(schemaConfig)
    if err != nil {
        fmt.Println("error: ", err)
    }

    // query execution
    query := `{
        hello
    }`

    params := graphql.Params{Schema: schema, RequestString: query}
    r := graphql.Do(params)

    if len(r.Errors) > 0 {
        fmt.Println("errors: ", r.Errors)
        return
    }

    rJSON, _ := json.Marshal(r)
    if err != nil {
        fmt.Println("error: ", err)
        return
    }

    fmt.Printf("%s \n", rJSON)
}
