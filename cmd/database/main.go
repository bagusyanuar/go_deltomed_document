package main

import (
	"flag"
	"log"

	"github.com/bagusyanuar/go_deltomed_document/cmd/database/scheme"
	"github.com/bagusyanuar/go_deltomed_document/cmd/database/seeder"
	"github.com/bagusyanuar/go_deltomed_document/config"
)

func main() {
	configuration := config.New()
	database := config.NewDatabaseConnection(configuration)
	seed := flag.String("m", "", "unsupport command")
	flag.Parse()
	command := *seed
	switch command {
	case "seed":
		seeder.Seed(database)
		return
	case "migrate":
		scheme.Migrate(database)
		return
	default:
		log.Println("unknown command")
		return
	}
}
