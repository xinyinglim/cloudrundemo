package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World\n")
}

func main() {
	portNo := os.Getenv("PORT")
	if portNo == "" {
		portNo = "8080"
	}
	http.HandleFunc("/", handler)
	fmt.Printf("Listening to port %s\n", portNo)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", portNo), nil))
}
