package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("New account Number :", acc.Number)
	return acc

}

func seedAccounts(s Storage) {
	seedAccount(s, "Diku", "GF", "Diku456")
}

func main() {

	seed := flag.Bool("seed", false, "Seed the DB")
	flag.Parse()

	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("Seeding the database")
		//Seed stuffs
		seedAccounts(store)
	}

	server := newApiServer(":4500", store)
	server.Run()
}
