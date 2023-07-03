package main

import (
	"fmt"
	"log"

	"github.com/s-antoshkin/go-comments-rest-api/internal/comment"
	"github.com/s-antoshkin/go-comments-rest-api/internal/db"
	transportHttp "github.com/s-antoshkin/go-comments-rest-api/internal/transport/http"
)

// Run - going to be responsible for
// the instantiation and startup of
// go applications
func Run() error {
	fmt.Println("starting up application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err = db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Comments REST API")
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
