package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", store)
	// server := newApiServer(":4500", store)
	// server.Run()
	fmt.Println("Shri Ganeshay Namh!")
}
