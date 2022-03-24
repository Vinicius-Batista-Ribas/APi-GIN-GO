package main

import (
	"github.com/viniciusRibas/api-go-gin/database"
	"github.com/viniciusRibas/api-go-gin/routes"
)

func main() {
	database.Conceta_BD()
	routes.Handle()
}
