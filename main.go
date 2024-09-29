package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ranjanvivesh/topmedia/router"
)

func main() {
	fmt.Println("Starting the Server......")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("Listening to port 5000")
}