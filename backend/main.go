package main

import (
	"fmt"
	"log"
	"net/http"
)

// DO STUFF HERE func doStuff() 


func handleRequests() {
	http.HandleFunc('/', boinc)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func main() {
	handleRequests()
}
