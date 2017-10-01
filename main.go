package main

import (
	"log"
	"os"

	"github.com/LevInteractive/allwrite-docs/api"
	"github.com/LevInteractive/allwrite-docs/gdrive"
	"github.com/LevInteractive/allwrite-docs/store/postgres"
	"github.com/LevInteractive/allwrite-docs/util"
	"github.com/joeshaw/envdecode"
	"github.com/urfave/cli"
)

var configMessage = `

#!/bin/bash

# These are the environmental variables you should set before running allwrite.
# You can set these in something like Upstart or in your current shell by doing:

# source ./creds.sh
# ------------------------------------------------------------------------------

# Basically, you just need to make sure these get set somehow, somewhere.

# The ID of the base directory for the docs (you can grab it from the URL in
# Drive).
export ACTIVE_DIR="xxxxxxxxxxxxxxxxxxx"

# Path to your Google client secret json file.
export CLIENT_SECRET="$PWD/client_secret.json"

# The storage system to use - currently postgres is the only option.
export STORAGE="postgres"
export PG_USER="root"
export PG_DB="allwrite"
export PG_HOST="localhost"

# Specify the port to run the application on.
export PORT=":8000"

# How often Google is queried for updates specified in milliseconds.
export FREQUENCY="300000"

`

var debugInfoMessage = `

------------------------------------------------------
INFO:
Client Secret: %s
Active Directory: %s
Storage Drive: %s
Address: %s
------------------------------------------------------

`

func main() {
	app := cli.NewApp()
	app.Name = "Allwrite Docs | Publish your documentation with Drive."
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Start the server in the foreground. This will authenticate with Google if it's the first time you're running.",
			Action: func(c *cli.Context) error {
				if env := setupConf(); env != nil {
					client := gdrive.DriveClient(env.CFG.ClientSecret)
					gdrive.WatchDrive(client, env)
					api.Listen(env)
				}
				return nil
			},
		},
		{
			Name:  "setup",
			Usage: "Only authenticate with Google and do not run the allwrite server.",
			Action: func(c *cli.Context) error {
				if env := setupConf(); env != nil {
					gdrive.DriveClient(env.CFG.ClientSecret)
				}
				return nil
			},
		},
		{
			Name:    "pull",
			Aliases: []string{"p"},
			Usage:   "Pull the latest content from Google Drive.",
			Action: func(c *cli.Context) error {
				if env := setupConf(); env != nil {
					client := gdrive.DriveClient(env.CFG.ClientSecret)
					gdrive.UpdateMenu(client, env)
				}
				return nil
			},
		},
		{
			Name:    "reset",
			Aliases: []string{"r"},
			Usage:   "Reset any saved authentication credentials for Google. You will need to re-authenticate after doing this.",
			Action: func(c *cli.Context) error {
				if env := setupConf(); env != nil {
					if err := gdrive.RemoveCacheFile(); err != nil {
						return err
					}
				}
				return nil
			},
		},
		{
			Name:    "info",
			Aliases: []string{"i"},
			Usage:   "Display environmental variables. Useful for making sure everything is setup correctly.",
			Action: func(c *cli.Context) error {
				if env := setupConf(); env != nil {
					log.Printf(
						debugInfoMessage,
						env.CFG.ClientSecret,
						env.CFG.ActiveDir,
						env.CFG.StoreType,
						env.CFG.Port,
					)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}

func setupConf() *util.Env {
	var cfg util.Conf
	if err := envdecode.Decode(&cfg); err != nil {
		log.Println(configMessage + "\n")
		return nil
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
			log.Printf("Could not connect to postgres: %s", err.Error())
			return nil
		}
	default:
		log.Printf("You must specify a storage system. (postgres)")
		return nil
	}

	return env
}
