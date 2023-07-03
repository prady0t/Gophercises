package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/home", home)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", )
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "At home")
}