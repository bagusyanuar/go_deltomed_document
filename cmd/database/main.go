package main

import (
	"github.com/bagusyanuar/go_deltomed_document/cmd/database/scheme"
	"github.com/bagusyanuar/go_deltomed_document/config"
)

func main() {
	configuration := config.New()
	database := config.NewDatabaseConnection(configuration)
	scheme.Migrate(database)
}
