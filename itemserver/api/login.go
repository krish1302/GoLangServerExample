package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)


func (ls *Server) loginRoutes(){
	ls.HandleFunc("/login/user",ls.createUserLogin()).Methods("POST")
	ls.HandleFunc("/login/users",ls.getUsersLogin()).Methods("GET")
	ls.HandleFunc("/login/user/{id}",ls.getUserLogin()).Methods("GET")
	ls.HandleFunc("/login/user/{id}",ls.deleteUserLogin()).Methods("DELETE")
	ls.HandleFunc("/login/user/{id}",ls.updateUserLogin()).Methods("PUT")
	ls.HandleFunc("/login/user/{id}",ls.patchUserLogin()).Methods("PATCH")
}

func (ls *Server) loginData(){
	ls.loginUsers =append(ls.loginUsers, Login{ID:"1",Name:"bala",Pass:"bala123"})
	ls.loginUsers =append(ls.loginUsers, Login{ID:"2",Name:"morris",Pass:"morris123"})
	ls.loginUsers =append(ls.loginUsers, Login{ID:"3",Name:"maha",Pass:"maha123"})
}

//creat new user 
func (ls *Server) createUserLogin() http.HandlerFunc{
	return func(rw http.ResponseWriter, r *http.Request) {
		var i Login
		if err:= json.NewDecoder(r.Body).Decode(&i); err != nil{
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		ls.loginUsers = append(ls.loginUsers, i)

		if err:= json.NewEncoder(rw).Encode(ls.loginUsers); err != nil{
			http.Error(rw,err.Error(),http.StatusInternalServerError)
			return
		}
	}
}

// get all user
func (ls *Server) getUsersLogin() http.HandlerFunc{
	return func(rw http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(rw).Encode(ls.loginUsers); err != nil{
			http.Error(rw,err.Error(),http.StatusInternalServerError)
			return
		}
	}
}

// get single user
// (obj *struct_name) access the object inside the struct defined
func (ls *Server) getUserLogin() http.HandlerFunc{
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r) // Gets params
	// Loop through users and find one with the id from the params
		for _, item := range ls.loginUsers {
			if item.ID == params["id"] {
				json.NewEncoder(rw).Encode(item)
				return
			}
		}
		json.NewEncoder(rw).Encode(&Login{})
	}
}
//delete user
func (ls *Server) deleteUserLogin() http.HandlerFunc{
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r) // Gets params
		for i, item := range ls.loginUsers {
			if item.ID == params["id"] {
				ls.loginUsers =append(ls.loginUsers[:i],ls.loginUsers[i+1:]... )
				json.NewEncoder(rw).Encode(ls.loginUsers)
				return
			}
		}
	}
}
//update user
func (ls *Server) updateUserLogin() http.HandlerFunc{
	return func(rw http.ResponseWriter, r *http.Request) {
		var user Login
		if err:= json.NewDecoder(r.Body).Decode(&user); err != nil{
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		params := mux.Vars(r) // Gets params
		for i, item := range ls.loginUsers {
			if item.ID == params["id"]{
				ls.loginUsers =append(ls.loginUsers[:i],ls.loginUsers[i+1:]... )
				ls.loginUsers =append(ls.loginUsers, user)
				json.NewEncoder(rw).Encode(ls.loginUsers)
				return
			}
		}
	}
}

//patch user
func (ls *Server) patchUserLogin() http.HandlerFunc{
	return func(rw http.ResponseWriter, r *http.Request) {
		
		params := mux.Vars(r) // Gets params
		for i, item := range ls.loginUsers {
			if item.ID == params["id"]{
				json.NewDecoder(r.Body).Decode(&ls.loginUsers[i])
				json.NewEncoder(rw).Encode(ls.loginUsers)
				return
			}
		}
	}
}