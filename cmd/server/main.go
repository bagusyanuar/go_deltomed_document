package main

import (
	"fmt"

	"github.com/bagusyanuar/go_deltomed_document/config"
	"github.com/bagusyanuar/go_deltomed_document/router"
)

func main() {
	configuration := config.New()
	database := config.NewDatabaseConnection(configuration)
	appPort := configuration.Get("APP_PORT")
	server := router.Build(database)
	server.Run(fmt.Sprintf(":%s", appPort))
}