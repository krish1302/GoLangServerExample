package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)



func (s *Server) shopRoutes(){
	s.HandleFunc("/shopping/item",s.createShopingItem()).Methods("POST")
	s.HandleFunc("/shopping/items",s.getShopingItems()).Methods("GET")
	s.HandleFunc("/shopping/items/{id}",s.getShopingItem()).Methods("GET")
	s.HandleFunc("/shopping/items/{id}",s.deleteShopingItem()).Methods("DELETE")
	s.HandleFunc("/shopping/items/{id}",s.updateShopingItem()).Methods("PUT")
	s.HandleFunc("/shopping/items/{id}",s.patchShopingItem()).Methods("PATCH")
}

func (s *Server) shopData(){
	s.shopingItems = append(s.shopingItems, Item{ID:"1",Product: "soap",Price: "10"})
	s.shopingItems = append(s.shopingItems, Item{ID:"2",Product: "plate",Price: "20"})
	s.shopingItems = append(s.shopingItems, Item{ID:"3",Product: "brush",Price: "30"})
}

func (s *Server) createShopingItem() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var i Item

		if err := json.NewDecoder(r.Body).Decode(&i); err != nil{
			http.Error(rw,err.Error(),http.StatusBadRequest)
			return
		}

		s.shopingItems = append(s.shopingItems, i)
		rw.Header().Set("Content-Type","application/json")
		if err := json.NewEncoder(rw).Encode(i); err != nil{
			http.Error(rw,err.Error(),http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) getShopingItems() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type","application/json")
		if err := json.NewEncoder(rw).Encode(s.shopingItems); err != nil{
			http.Error(rw,err.Error(),http.StatusInternalServerError)
			return
		}
	}
}

// get single item
// (obj *struct_name) access the object inside the struct defined
func (ls *Server) getShopingItem() http.HandlerFunc{
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type","application/json")
		params := mux.Vars(r) // Gets params
	// Loop through users and find one with the id from the params
		for _, item := range ls.shopingItems {
			if item.ID == params["id"] {
				json.NewEncoder(rw).Encode(item)
				return
			}
		}
		json.NewEncoder(rw).Encode(&Login{})
	}
}
//delete user
func (ls *Server) deleteShopingItem() http.HandlerFunc{
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin","*")
		rw.Header().Set("Content-Type","application/json")

		params := mux.Vars(r) // Gets params
		for i, item := range ls.shopingItems {
			if item.ID == params["id"] {
				ls.shopingItems =append(ls.shopingItems[:i],ls.shopingItems[i+1:]... )
				json.NewEncoder(rw).Encode(ls.shopingItems)
				return
			}
		}
	}
}
//update user
func (ls *Server) updateShopingItem() http.HandlerFunc{
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type","application/json")
		var product Item
		if err:= json.NewDecoder(r.Body).Decode(&product); err != nil{
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		params := mux.Vars(r) // Gets params
		for i, item := range ls.shopingItems {
			if item.ID == params["id"]{
				ls.shopingItems =append(ls.shopingItems[:i],ls.shopingItems[i+1:]... )
				ls.shopingItems =append(ls.shopingItems, product)
				json.NewEncoder(rw).Encode(ls.shopingItems)
				return
			}
		}
	}
}

//patch user
func (ls *Server) patchShopingItem() http.HandlerFunc{
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type","application/json")
		params := mux.Vars(r) // Gets params
		for i, item := range ls.shopingItems {
			if item.ID == params["id"]{
				json.NewDecoder(r.Body).Decode(&ls.shopingItems[i])
				json.NewEncoder(rw).Encode(ls.shopingItems)
				return
			}
		}
	}
}