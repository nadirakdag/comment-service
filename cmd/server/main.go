package main

import (
	"comment-service/config"
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

	return nil
}

func initConfig() {
	err := envconfig.Process("", &cfg)
	if err != nil {
		fmt.Errorf(err.Error())
	}
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Errorf("error whil starting application, err : %v", err)
	}
}
