package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"scm.bluebeam.com/stu/golang-template/apis/converters"
	"scm.bluebeam.com/stu/golang-template/apis/dtos"
	"scm.bluebeam.com/stu/golang-template/errors"
	"scm.bluebeam.com/stu/golang-template/services"
)

type (
	// UsersResourcer provides an inverface for user resources
	UsersResourcer interface {
		GetUser(res http.ResponseWriter, req *http.Request)
		CreateUser(res http.ResponseWriter, req *http.Request)
	}

	// UsersResource defines handlers for the APIs
	UsersResource struct {
		Service services.UsersServicer
	}
)

// RegisterUsersResource sets up the routing of users endpoints and handlers
func RegisterUsersResource(router *mux.Router, service services.UsersServicer) {
	r := &UsersResource{service}
	router.HandleFunc("/users/{userID}", r.GetUser).Methods("GET")
	router.HandleFunc("/users", r.CreateUser).Methods("POST")
}

// GetUser by ID
func (r *UsersResource) GetUser(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	params := mux.Vars(req)
	userIDString := params["userID"]
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		http.Error(res, "UserID must be an integer", http.StatusBadRequest)
		return
	}
	serviceUser, err := r.Service.GetUser(userID)
	if err != nil {
		if _, ok := err.(errors.NotFound); ok {
			http.Error(res, fmt.Sprintf("User with ID %d not found", userID), http.StatusNotFound)
		} else {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	user := converters.ToUser(serviceUser)
	json.NewEncoder(res).Encode(user)
}

// CreateUser and return result
func (r *UsersResource) CreateUser(res http.ResponseWriter, req *http.Request) {
	var user dtos.User
	defer req.Body.Close()
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	serviceUser, err := r.Service.CreateUser(converters.FromUser(&user))
	if err != nil {
		if _, ok := err.(errors.InvalidArgument); ok {
			http.Error(res, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	resultUser := converters.ToUser(serviceUser)
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(resultUser)
}
