package api

import "github.com/gorilla/mux"

type Server struct {
	*mux.Router
	shopingItems []Item
	loginUsers   []Login
}

func NewServer() *Server {
	s := &Server{
		Router:       mux.NewRouter(),
		shopingItems: []Item{},
		loginUsers:   []Login{},
	}
	s.shopRoutes()
	s.loginRoutes()
	s.loginData()
	s.shopData()
	return s
}