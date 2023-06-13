package main

import (
    "fmt"
	"log" 
	"net/http"

)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Db Connected !")
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main () {

	handleRequests()

}


