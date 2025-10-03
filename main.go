package main

import (
	"go-api/database"
	"go-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequestes()
}
