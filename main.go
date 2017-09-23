package main

import (
	"errors"
	"log"

	"github.com/LevInteractive/allwrite-docs/gdrive"
	"github.com/LevInteractive/allwrite-docs/store/postgres"
	"github.com/LevInteractive/allwrite-docs/util"
	"github.com/joeshaw/envdecode"
)

func main() {
	var cfg util.Conf

	if err := envdecode.Decode(&cfg); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	env := &util.Env{
		CFG: &cfg,
	}

	switch cfg.StoreType {
	case "postgres":
		if db, err := postgres.Init(
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresDBName,
		); err == nil {
			env.DB = db
		} else {
			panic(err)
		}
	default:
		panic(errors.New("you must specify a storage system"))
	}

	gdrive.UpdateMenu(env)
}
