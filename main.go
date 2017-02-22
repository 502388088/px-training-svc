package main

import (
	"github.com/502388088/px-training-svc/server"
	"github.com/502388088/px-training-svc/postgres"
)

func main() {
	postgres.Open_postgres("user=postgres password=admin dbname=postgres sslmode=disable", "people")
	defer postgres.Database.Close()

	server.SetupServer()
	server.StartRestServer("8089")
}
