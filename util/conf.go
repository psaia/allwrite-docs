package util

// Conf - Configuration.
type Conf struct {
	ActiveDir string `env:"ACTIVE_DIR,required"`
	RedisHost string `env:"REDIS_HOST,default=localhost"`
	RedisPort string `env:"REDIS_PORT,default=6379"`
}
