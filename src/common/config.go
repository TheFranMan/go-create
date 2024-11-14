package common

import (
	"strings"

	"github.com/caarlos0/env"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

type Envs struct {
	IsLocal bool
	IsStage bool
	IsProd  bool
}

type Config struct {
	Envs
	Port int    `env:"PORT"`
	Env  string `env:"ENV"`
}

func GetConfig() (*Config, error) {
	var config Config
	err := env.Parse(&config)
	if nil != err {
		return nil, err
	}

	config.setEnv()
	config.setLogging()

	return &config, nil
}

func (c *Config) setEnv() {
	if strings.HasPrefix(strings.ToLower(c.Env), "prod") {
		c.IsLocal = false
		c.IsStage = false
		c.IsProd = true
		return
	}

	if strings.HasPrefix(strings.ToLower(c.Env), "stag") {
		c.IsLocal = false
		c.IsStage = true
		c.IsProd = false
		return
	}

	c.IsLocal = true
	c.IsStage = false
	c.IsProd = false
}

func (c *Config) setLogging() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.SetLevel(log.InfoLevel)

	if c.IsLocal {
		log.SetLevel(log.DebugLevel)
	}
}
