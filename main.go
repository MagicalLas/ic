package main

import (
	"fmt"
	"gomod.usaken.org/ic/api"
	"gomod.usaken.org/ic/config"
	"gomod.usaken.org/ic/monitor"
	"gomod.usaken.org/ic/spine"
	"log"
	"os"
)

func main() {
	fmt.Println("impression counter system started.")

	spine.CheckNetworkInit()

	c, err := config.Load()
	if err != nil {
		log.Fatalf("config load failed. %e\n", err)
	}

	monitor.RunPprofServer(c)

	err = monitor.RunPrometheusServer(c)
	if err != nil {
		log.Fatalf("prometheus metric server fail to start. %e\n", err)
	}

	err = api.Run(c)
	if err != nil {
		log.Fatalf("server fail to start. %e\n", err)
	}

	spine.WaitUntilSystemShutdown()

	fmt.Println("impression counter system ended.")
	os.Exit(0)
}
