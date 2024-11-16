package config

import (
	"fmt"
	"gomod.usaken.org/ic/spine"
	"runtime"
)

func Load() (*Config, error) {
	fmt.Printf("config load start\n")
	spine.SystemGroup.Add(1)
	defer spine.SystemGroup.Done()

	c := &Config{
		ServerAddr:             ":8081",
		ServerName:             "ic",
		Concurrency:            2,
		EnableProfiling:        true,
		EnablePrometheusMetric: true,
	}

	maxprocs := runtime.GOMAXPROCS(-1)
	runtime.GOMAXPROCS(c.Concurrency)
	newmaxprocs := runtime.GOMAXPROCS(-1)

	cores := runtime.NumCPU()
	fmt.Printf("(- DEFAULT:GOMAXPROCS=%v -)\n", maxprocs)
	fmt.Printf("(- NOW:GOMAXPROCS=%v -)\n", newmaxprocs)

	fmt.Printf("(- core=%v -)\n", cores)
	fmt.Printf("(- PORT=%v -)\n", c.ServerAddr)

	fmt.Printf("config loaded\n")
	return c, nil
}

type Config struct {
	ServerAddr             string
	ServerName             string
	Concurrency            int
	EnableProfiling        bool
	EnablePrometheusMetric bool
}
