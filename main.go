// package main

// import "fmt"

// func main() {

// 	fmt.Println("Brian")

// 	var name string
// 	name = "Brian"

// 	var studentName = "Lauren"
// 	studentName = "Bagus"

// 	batchName := "Batch 40"

// 	var bialngan1 = 20
// 	var bilangan2 = 50

// 	fmt.Println(bialngan1 + bilangan2)

// 	fmt.Println(studentName)
// 	fmt.Println(name)
// 	fmt.Println(batchName)

// }
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).Methods("GET")

	route.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Contact Me"))
	}).Methods("GET")

	fmt.Println("Server running on port 5000")
	http.ListenAndServe("localhost:5000", route)

}
