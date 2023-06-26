package main

import (
	"fmt"
	"log"
)

// Run - going to be responsible for
// the instantiation and startup of
// go applications
func Run() error {
	fmt.Println("starting up application")
	return nil
}

func main() {
	fmt.Println("Comments REST API")
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
