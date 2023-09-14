package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/androide18/go-user-api/pkg/config"
	"github.com/androide18/go-user-api/pkg/models"
	"github.com/androide18/go-user-api/pkg/services"
	"github.com/androide18/go-user-api/pkg/utils"
	"github.com/gorilla/mux"
)

var NewUser models.User

var userService *services.UserService

func init() {
	userService = services.NewUserService(config.GetDB())
}

// Modify the controller functions to use userService functions
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := userService.GetAllUsers()
	res, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	createUser := &models.User{}
	utils.ParseBody(r, createUser)
	u := userService.CreateUser(createUser)
	res, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// With channel comming from service
func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		http.Error(w, "Error while parsing", http.StatusBadRequest)
		return
	}

	// Use the service method with the channel
	userChan := userService.GetUserByIdWithChannel(ID)
	userDetails, ok := <-userChan
	if !ok {
		http.Error(w, "Error fetching user details", http.StatusInternalServerError)
		return
	}

	if userDetails == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(userDetails)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	user := userService.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// with channel from controller
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 10, 64) // Updated parsing base to 10 and size to 64
	if err != nil {
		http.Error(w, "Error while parsing", http.StatusBadRequest)
		return
	}

	// Get the user details using the service
	userDetails, err := userService.GetUserById(ID)
	if err != nil {
		http.Error(w, "Error fetching user details", http.StatusInternalServerError)
		return
	}

	// Use a WaitGroup to wait for concurrent tasks to complete
	var wg sync.WaitGroup

	// Create a channel to send updates result
	updateResultChan := make(chan *models.User)

	// Update user details concurrently if needed
	if updateUser.Name != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			userDetails.Name = updateUser.Name
		}()
	}
	if updateUser.Lastname != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			userDetails.Lastname = updateUser.Lastname
		}()
	}
	if updateUser.Email != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			userDetails.Email = updateUser.Email
		}()
	}

	// Wait for all concurrent tasks to complete
	go func() {
		wg.Wait()
		updateResultChan <- userDetails
		close(updateResultChan)
	}()

	updatedUser := <-updateResultChan

	// Save the updated user details using the service
	updatedUser, err = userService.UpdateUser(updatedUser)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(updatedUser)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
