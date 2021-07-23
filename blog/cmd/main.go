package main

import (
	"fmt"
	"log"
	"net/http"

	"blog/pkg/adding"
	"blog/pkg/http/rest"
	"blog/pkg/listing"
	"blog/pkg/storage/json"
)

func main() {
	fmt.Println("blog started...")

	var adder adding.Service
	var lister listing.Service

	// error handling omitted for simplicity
	s, _ := json.NewStorage()

	adder = adding.NewService(s)
	lister = listing.NewService(s)

	// add default data
	// adder.AddDefaultPosts(DefaultPosts)
	// fmt.Println("Finished adding sample data.")

	router := rest.Handler(adder, lister)
	fmt.Println("blog now at service (localhost:8080)...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
