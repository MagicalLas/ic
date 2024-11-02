package config

import (
	"fmt"
	"gomod.usaken.org/ic/spine"
)

func Load() (*Config, error) {
	fmt.Printf("config load start\n")
	spine.SystemGroup.Add(1)
	defer spine.SystemGroup.Done()

	fmt.Printf("config loaded\n")
	return &Config{
		ServerAddr: ":8081",
		ServerName: "ic",
	}, nil
}

type Config struct {
	ServerAddr string
	ServerName string
}
