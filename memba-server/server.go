package main

import (
	"fmt"
	"log"
	"os"

	"github.com/antiharmonic/memba/memba-server/memba"
)

func main() {
	// configuration
	log.SetOutput(os.Stdout)
	var config memba.Config
	err := memba.LoadConfig(&config)

	if err != nil {
		log.Fatalln(err)
	}

	// database

	// set up handlers

	// start
}

func DataHandler() {
	// get previous, random current, random next
	fmt.Println("test")
}

// TODO: set up integrations, users
