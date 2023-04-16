package main

import (
	"comment-service/config"
	"comment-service/internal/comment"
	"comment-service/internal/db/postgres"
	transportHttp "comment-service/internal/transport/http"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

var cfg config.Config

// Run - is going to be responsible for
// the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("starting up our application")

	initConfig()

	db, err := postgres.NewDatabase(cfg.Database)
	if err != nil {
		fmt.Printf("failed to connect to the database, err: %v \n", err)
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println(err)
		return err
	}

	commentService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(commentService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func initConfig() {
	err := envconfig.Process("", &cfg)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Printf("error while starting application, err : %v \n", err)
	}
}
