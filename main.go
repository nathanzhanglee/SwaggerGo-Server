package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// User represents the user for this application
// swagger:model
type User struct {
	// The name for this user
	// required: true
	// minLength: 5
	Name string `json:"name"`

	// The birth year for this user
	// minimum: 1900
	// maximum: 2022
	BirthYear int `json:"birth_year"`
}

func main() {
	router := mux.NewRouter()

	// Get Users
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// swagger:operation GET /users getUsers
		//
		// Retrieves all users
		//
		// ---
		// produces:
		//   - application/json
		// responses:
		//   '200':
		//     description: A list of users
		//     schema:
		//       type: array
		//       items:
		//         "$ref": "#/definitions/User"

		users := []User{
			{"Mario", 1990},
			{"Wario", 1980},
		}

		res, _ := json.Marshal(&users)

		_, _ = w.Write(res)
	}).Methods(http.MethodGet)

	// Create User
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// swagger:operation POST /users postUser
		//
		// Create a new user
		//
		// ---
		// produces:
		//   - application/json
		// parameters:
		//   - name: body
		//     in: body
		//     description: User object to create
		//     required: true
		//     schema:
		//       "$ref": "#/definitions/User"
		// responses:
		//   '200':
		//     description: User successfully created
		//     schema:
		//       "$ref": "#/definitions/User"
		//   '400':
		//     description: Invalid input
		//   '500':
		//     description: Internal server error

		var user User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			_, _ = w.Write([]byte("decoding failed"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, _ = w.Write([]byte(fmt.Sprintf("created %+v", user)))
	}).Methods(http.MethodPost)

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(s.ListenAndServe())
}
