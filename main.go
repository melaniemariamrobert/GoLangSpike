// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

// These are the libraries we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (
	// "fmt" has methods for formatted I/O operations (like printing to the console)
	"fmt"
	// The "net/http" library has methods to implement HTTP clients and servers
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// Declare a new router
    	r := mux.NewRouter()

    	// This is where the router is useful, it allows us to declare methods that
    	// this path will be valid for
    	r.HandleFunc("/hello", handler).Methods("GET")

    	// We can then pass our router (after declaring all our routes) to this method
    	// (where previously, we were leaving the secodn argument as nil)
    	http.ListenAndServe(":8080", r)
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello Melanie!")
}