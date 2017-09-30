package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LevInteractive/allwrite-docs/api"
	"github.com/LevInteractive/allwrite-docs/gdrive"
	"github.com/LevInteractive/allwrite-docs/store/postgres"
	"github.com/LevInteractive/allwrite-docs/util"
	"github.com/joeshaw/envdecode"
	"github.com/urfave/cli"
)

func main() {

	// Setup and validate environmental variables.
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
			fmt.Printf("Could not connect to postgres: %s", err.Error())
			return
		}
	default:
		fmt.Printf("you must specify a storage system. (postgres)")
		return
	}

	// Parse CLI commands.
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Start the server in the foreground. This will authenticate with Google if it's the first time you're running.",
			Action: func(c *cli.Context) error {
				gdrive.DriveClient(env.CFG.ClientSecret)
				api.Listen(env)
				return nil
			},
		},
		{
			Name:    "reset",
			Aliases: []string{"r"},
			Usage:   "Reset any saved authentication credentials for Google. You will need to re-authenticate after doing this.",
			Action: func(c *cli.Context) error {
				if err := gdrive.RemoveCacheFile(); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "setup",
			Usage: "Only authenticate with Google and do not run the allwrite server.",
			Action: func(c *cli.Context) error {
				gdrive.DriveClient(env.CFG.ClientSecret)
				return nil
			},
		},
		{
			Name:    "pull",
			Aliases: []string{"p"},
			Usage:   "Pull the latest content from Google Drive.",
			Action: func(c *cli.Context) error {
				client := gdrive.DriveClient(env.CFG.ClientSecret)
				if err := gdrive.UpdateMenu(client, env); err != nil {
					fmt.Printf("Could not pull the latest from Google: %s", err.Error())
				}
				return nil
			},
		},
		{
			Name:    "info",
			Aliases: []string{"i"},
			Usage:   "Display environmental variables. Useful for making sure everything is setup correctly.",
			Action: func(c *cli.Context) error {
				fmt.Printf(`
------------------------------------------------------
INFO:
Client Secret: %s
Active Directory: %s
Storage Drive: %s
Address: %s
------------------------------------------------------\n
				`, cfg.ClientSecret, cfg.ActiveDir, cfg.StoreType, cfg.Port)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
