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
	app.Version = "0.0.2"
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
		log.Println("Please make sure the environmental variables are set first.")
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
