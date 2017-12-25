package util

import (
	"github.com/LevInteractive/allwrite-docs/store"
)

// Conf - Configuration.
type Conf struct {
	ActiveDir        string `env:"ACTIVE_DIR,required"`
	StoreType        string `env:"STORAGE,required"`
	ClientSecret     string `env:"CLIENT_SECRET,required"`
	Frequency        int    `env:"FREQUENCY,required"`
	Port             string `env:"PORT,required"`
	CertbotEmail     string `env:"CERTBOT_EMAIL"`
	Domain           string `env:"DOMAIN"`
	PostgresUser     string `env:"PG_USER"`
	PostgresPassword string `env:"PG_PASS"`
	PostgresHost     string `env:"PG_HOST"`
	PostgresDBName   string `env:"PG_DB"`
}

// Env is a environment specific struct.
type Env struct {
	DB  store.Store
	CFG *Conf
}
