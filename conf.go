package main

import (
	"fmt"
	"os"
)

type Conf struct {
	Host string
	Port string
}

func NewConf() *Conf {
	conf := &Conf{
		Host: "127.0.0.1",
		Port: "8000",
	}
	conf.load()
	return conf
}

func (c *Conf) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func (c *Conf) load() {
	if host, ok := os.LookupEnv("HOST"); ok {
		c.Host = host
	}
	if port, ok := os.LookupEnv("PORT"); ok {
		c.Port = port
	}
}
