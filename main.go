package main

import (
	"fmt"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/srgrcp/Harper-Test/resolver"
	"github.com/srgrcp/Harper-Test/schema"
)

//go:generate go run generator/main.go

func main() {
	schema := graphql.MustParseSchema(schema.Data, resolver.NewRoot())
	http.Handle("/graphql", &relay.Handler{Schema: schema})
	http.Handle("/graphql/", &relay.Handler{Schema: schema})
	fmt.Println("http://localhost:8080/graphql")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
