package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User represents a user in the system.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// users is a slice to store users.
var Users = []User{
	{ID: 1, Name: "Foo Bar", Email: "foobar@example.com"},
	{ID: 2, Name: "Fizz Buzz", Email: "fizzbuzz@example.com"},
}
var nextId = 3

// GetUserByID handles the GET /users/{id} endpoint.
//   - HTTP Method: GET
//   - Description: Retrieves the details of a specific user.
//   - Parameters: id (path parameter, integer, required) - The ID of the user to retrieve.
//   - Response Format: A JSON object representing the user. If no user with the given ID is found, a 404 error is returned with the message "No user found!".
//   - Example: curl http://localhost:3000/users/1
//   - Test: http://localhost:3000/get
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for _, user := range Users {
		if user.ID == id { // changed from User.ID to user.ID
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "No user found!", http.StatusNotFound)
}

// CreateUser handles the POST /users endpoint.
//   - HTTP Method: POST
//   - Description: Creates a new user.
//   - Parameters: User (body parameter, JSON object, required) - A JSON object representing the user to create. This object should include Name and Email fields.
//   - Response Format: A JSON object representing the created user.
//   - Example: curl -X POST -d '{"Name":"John Doe","Email":"johndoe@example.com"}' http://localhost:3000/users
//   - Test: http://localhost:3000/post
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = nextId // generate a new ID for the user
	Users = append(Users, user)
	json.NewEncoder(w).Encode(user)
	nextId++
}

// EditUser handles the PUT /users/{id} endpoint.
//   - HTTP Method: PUT
//   - Description: Edits the details of a specific user.
//   - Parameters: id (path parameter, integer, required) - The ID of the user to edit. User (body parameter, JSON object, required) - A JSON object representing the new details of the user. This object should include Name and Email fields.
//   - Response Format: A JSON object representing the updated user. If no user with the given ID is found, a 404 error is returned with the message "No user found!".
//   - Example: curl -X PUT -d '{"Name":"John Doe","Email":"johndoe@example.com"}' http://localhost:3000/users/1
//   - Test: http://localhost:3000/put
func EditUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var updatedUser User
	_ = json.NewDecoder(r.Body).Decode(&updatedUser)
	for i, user := range Users {
		if user.ID == id {
			Users[i].Name = updatedUser.Name
			Users[i].Email = updatedUser.Email
			json.NewEncoder(w).Encode(Users[i])
			return
		}
	}
	http.Error(w, fmt.Sprintf("User ID %d not found", id), http.StatusNotFound)
}

// DeleteUser handles the DELETE /users/{id} endpoint.
//   - HTTP Method: DELETE
//   - Description: Deletes a specific user.
//   - Parameters: id (path parameter, integer, required) - The ID of the user to delete.
//   - Response Format: A success message indicating that the user has been deleted. If no user with the given ID is found, a 404 error is returned with the message "No user found!".
//   - Example: curl -X DELETE http://localhost:3000/users/1
//   - Test: http://localhost:3000/delete
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for i, user := range Users {
		if user.ID == id {
			Users = append(Users[:i], Users[i+1:]...)
			fmt.Fprintf(w, "User with ID %v has been deleted successfully", id)
			return
		}
	}
	http.Error(w, "No user found!", http.StatusNotFound)
}
