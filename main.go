package main

import (
	"log"

	"github.com/LevInteractive/allwrite-docs/gdrive"
	"github.com/LevInteractive/allwrite-docs/util"
	"github.com/joeshaw/envdecode"
)

func main() {
	var cfg util.Conf
	if err := envdecode.Decode(&cfg); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	service := gdrive.DriveClient()
	gdrive.UpdateMenu(&cfg, service)
}
