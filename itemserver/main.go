package main

import (
	"main/api"
	"net/http"

	"github.com/rs/cors"
)

func main(){
	server := api.NewServer()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
        AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPatch},
        AllowCredentials: true,
    }).Handler(server)
	
	http.ListenAndServe(":8888",c)
}