package main

import (
	"fmt"
	"os"
)

type Conf struct {
	Host string
	Port string
}

func (c *Conf) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

var conf *Conf

func init() {
	conf = &Conf{
		Host: "127.0.0.1",
		Port: "8000",
	}

	if host, ok := os.LookupEnv("HOST"); ok {
		conf.Host = host
	}
	if port, ok := os.LookupEnv("PORT"); ok {
		conf.Port = port
	}
}
