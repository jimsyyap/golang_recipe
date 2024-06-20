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
    author := &Author{Name: "jimmy desu", Tutorials: []int{1}}
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

    tutorials := populate()

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
        "tutorial": &graphql.Field{
            Type: tutorialType,
            Description: "Get Tutorial By ID",
            Args: graphql.FieldConfigArgument{
                "id": &graphql.ArgumentConfig{
                    Type: graphql.Int,
                },
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                id, ok := p.Args["id"].(int)
                if ok {
                    for _, tutorial := range tutorials {
                        if int(tutorial.ID) == id {
                            return tutorial, nil
                        }
                    }
                }
                return nil, nil
            },
        },
        "list": &graphql.Field{
            Type: graphql.NewList(tutorialType),
            Description: "Get Tutorial List",
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return tutorials, nil
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
        list {
            id
            title
        }
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