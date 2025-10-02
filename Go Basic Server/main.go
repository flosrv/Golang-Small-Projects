// package declaration
package main

// import necessary packages
import (
	"fmt"      // for printing output
	"log"      // for logging errors
	"net/http" // for HTTP server and request handling
)

// handler for GET /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// check if URL path is exactly /hello
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound) // respond 404 if path mismatch
		return
	}

	// check if HTTP method is GET
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed) // respond 405 if wrong method
		return
	}

	// send response "Hello!" to client
	fmt.Fprintf(w, "Hello!")
}

// handler for POST /form
func formHandler(w http.ResponseWriter, r *http.Request) {
	// check if HTTP method is POST
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // respond 405 if wrong method
		return
	}

	// parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() err: %v", err), http.StatusBadRequest) // respond 400 on error
		return
	}

	// get form values
	name := r.FormValue("name")   // retrieve "name" field
	email := r.FormValue("email") // retrieve "address" field

	// send response with submitted data
	fmt.Fprintf(w, "POST request successful\nName = %s\nemail = %s\n",
		name, email)
}

// main function - entry point of the program
func main() {
	// create a file server to serve static files from ./static directory
	fileServer := http.FileServer(http.Dir("./static"))

	// handle routes
	http.HandleFunc("/hello", helloHandler) // map /hello to helloHandler
	http.HandleFunc("/form", formHandler)   // map /form to formHandler
	http.Handle("/", fileServer)            // serve static files for root path and others

	// print message to console indicating server start
	fmt.Println("Starting server on :8080")

	// start HTTP server on port 8080 and log fatal error if server fails
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
