package main

import (
	"fmt"
	"os"

	"github.com/dkder3k/shg-srv/pkg/comm"
	"github.com/dkder3k/shg-srv/pkg/handler"
	"github.com/dkder3k/shg-srv/pkg/repository"
	"github.com/dkder3k/shg-srv/pkg/service"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
)

func main() {
	if err := initConfig(); err != nil {
		fmt.Printf("error during start: %s\n", err.Error())
		os.Exit(1)
	}

	consulClient, err := api.NewClient(&api.Config{
		Address:   "127.0.0.1:8500",
		Scheme:    "http",
		Transport: cleanhttp.DefaultPooledTransport(),
	})
	if err != nil {
		fmt.Printf("error during start: %s\n", err.Error())
		os.Exit(1)
	}

	repo := repository.NewConsulRepository(consulClient)
	services := service.NewService(repo)
	hndlr := handler.NewHandler(services)

	srv := new(comm.Server)
	if err := srv.Run("8080", hndlr.InitRoutes()); err != nil {
		fmt.Printf("error during start: %s\n", err.Error())
		os.Exit(1)
	}
}

func initConfig() error {
	return nil
}
