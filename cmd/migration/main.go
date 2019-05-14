package main

import (
	_ "github.com/lib/pq"
	"github.com/payfazz/fazzlearning-api/config"
	"github.com/payfazz/fazzlearning-api/database/migration"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

func main() {
	fazzdb.Migrate(config.GetDb(), "example", config.ForceMigrate(),
		migration.VersionExample,
	)
}
