package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/srgrcp/Harper-Test/api"
	"github.com/srgrcp/Harper-Test/orm"
	"github.com/srgrcp/Harper-Test/resolver"
	"github.com/srgrcp/Harper-Test/schema"
)

//go:generate go run generator/main.go

func main() {
	orm.InitDB()

	router := mux.NewRouter()
	router.HandleFunc("/api/order/{uuid}", api.GetServiceOrder).Methods("GET")

	schema := graphql.MustParseSchema(schema.Data, resolver.NewRoot())
	router.Handle("/graphql", &relay.Handler{Schema: schema}).Methods("POST")
	router.Handle("/graphql/", &relay.Handler{Schema: schema}).Methods("POST")
	/* http.Handle("/graphql", &relay.Handler{Schema: schema})
	http.Handle("/graphql/", &relay.Handler{Schema: schema}) */
	fmt.Println("http://localhost:8080/graphql")
	fmt.Println(http.ListenAndServe(":8080", router))
}
