package main

import (
	"fmt"
	"log"

	"gomod.usaken.org/ic/api"
	"gomod.usaken.org/ic/config"
	"gomod.usaken.org/ic/spine"
)

func main() {
	fmt.Println("impression counter system started.")

	spine.CheckNetworkInit()

	c, err := config.Load()
	if err != nil {
		log.Fatalf("config load failed. %e\n", err)
	}

	err = api.Run(c)
	if err != nil {
		log.Fatalf("server fail to start. %e\n", err)
	}

	spine.SystemGroup.Wait()

	fmt.Println("impression counter system ended.")
}
