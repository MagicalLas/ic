package monitor

import (
	"gomod.usaken.org/ic/config"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func RunPprofServer(c *config.Config) {
	if !c.EnableProfiling {
		return
	}

	runtime.SetBlockProfileRate(5)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}
