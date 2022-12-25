package config

import (
	"fmt"
	"os"
)

type Config struct {
	Host string
	Port string
}

func NewConfig() *Config {
	c := &Config{
		Host: "127.0.0.1",
		Port: "8000",
	}

	if host, ok := os.LookupEnv("HOST"); ok {
		c.Host = host
	}
	if port, ok := os.LookupEnv("PORT"); ok {
		c.Port = port
	}
	return c
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
