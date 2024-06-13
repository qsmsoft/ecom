package main

import (
	"fmt"
	"log"

	"github.com/qsmsoft/ecom/cmd/api"
	"github.com/qsmsoft/ecom/config"
	"github.com/qsmsoft/ecom/db"
)

func main() {

	cfg := db.Config{
		Host:     config.Envs.Host,
		Port:     config.Envs.Port,
		User:     config.Envs.User,
		Password: config.Envs.Password,
		DBName:   config.Envs.DBName,
		SSLMode:  config.Envs.SSLMode,
	}

	conn, err := db.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to the database!")

	server := api.NewAPIServer(":3000", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
