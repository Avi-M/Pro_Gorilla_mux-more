// calls.go
package main

// Imports
import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// Student ...
type StudentInfo struct {
	ID     int
	Name   string
	Class string
	Branch  string
}

// GetAllStudents ...
func GetAllStudents(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: GetAllStudents")

	// Post Request
	if r.Method == http.MethodPost {

		// Get the values from Form Data
		requestBody, err := json.Marshal(map[string]string{
			"name":   r.FormValue("name"),
			"class": r.FormValue("class"),
			"branch":  r.FormValue("branch"),
		})

		// Error Check
		if err != nil {
			fmt.Println("Error Occured")
			return
		}

		// Send the post request with the required Values
		resp, postErr := http.Post("http://localhost:8000/api/student/", "application/json", bytes.NewBuffer(requestBody))

		if postErr != nil {
			fmt.Println("Error Occured")
			return
		}

		fmt.Println(resp.Body)
	}

	// Fetch All the Student
	response, err := http.Get("http://localhost:8000/api/student")
	if err != nil {
		fmt.Printf("Could Not Fetch Foods, Error: %s", err)
		return
	}

	defer response.Body.Close()

	// Store the Fetched Student in an Array
	var items []StudentInfo
	_ = json.NewDecoder(response.Body).Decode(&items)

	// Template
	templ := template.Must(template.ParseFiles("templates/home.html"))

	templ.Execute(w, items)
	return
}

// updateStudentInfo ...
func updateStudentInfo(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: Updating Food Item")

	// Retrieve Student ID
	id := r.FormValue("id")

	// New Values
	requestBody, err := json.Marshal(map[string]string{
		"name":   r.FormValue("name"),
		"class": r.FormValue("class"),
		"branch":  r.FormValue("branch"),
	})

	// Error Check
	if err != nil {
		fmt.Println("Error Occured")
		return
	}

	// Specific URL for updating the required Item
	url := "http://localhost:8000/api/student/" + id + "/"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		fmt.Println("error occured")
		return
	}
	fmt.Println(resp.Body)

	// Redirect to Home Page
	http.Redirect(w, r, "/", 301)

}

func deleteStudentInfo(w http.ResponseWriter, r *http.Request) {

	// Send Post request to the URL with id to delete the Item
	id := r.FormValue("id")

	url := "http://localhost:8000/api/student/delete/" + id

	resp, err := http.Post(url, "application/json", nil)

	if err != nil {
		fmt.Println("error occured")
		return
	}

	fmt.Println(resp.Body)

	// Redirect to Home Page
	http.Redirect(w, r, "/", 301)
}

func main() {

	// Handler functions
	http.HandleFunc("/", GetAllStudents)
	http.HandleFunc("/update", updateStudentInfo)
	http.HandleFunc("/delete", deleteStudentInfo)
	http.ListenAndServe(":8080", nil)
}
