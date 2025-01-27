package main

import (
	"os"
	"strconv"
)

type Config struct {
	Port uint16 `json:"port" yaml:"port"`
}

func NewConfig() *Config {
	c := Config{
		Port: 8080,
	}

	if port, ok := os.LookupEnv("PORT"); ok {
		if p, err := strconv.ParseUint(port, 10, 16); err == nil {
			c.Port = uint16(p)
		}
	}

	return &c
}
