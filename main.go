package main

import (
	"log"

	"github.com/rezacse/go-micro/internal/database"
	"github.com/rezacse/go-micro/internal/server"
)

func main() {
	db, err := database.NewDtabaseClient()
	if err != nil {
		log.Fatalf("failed to init DB Client: %s", err)
	}

	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal((err.Error()))
	}

}