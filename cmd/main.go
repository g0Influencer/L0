package main

import (
	"L0/Nats"
	"L0/database"
	"L0/server"
)

func main() {

	database.Recover()
	Nats.Subscribe()
	server.Run()

}
