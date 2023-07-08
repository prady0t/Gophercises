package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
    HandleRequests()
	fmt.Println("started at: 8081")
}

func HelloHome(w http.ResponseWriter, r *http.Request){
	  // fmt.Fprintf(w, r.Response.Status)
      fmt.Fprintf(w, "Homepage HIt. Hello World")
}
func HandleRequests(){
	http.HandleFunc("/", HelloHome)
	log.Fatal(http.ListenAndServe(":8081", nil))
}