package scripts

import "log"

var stopScripts = make(chan struct{})

// Run - runs all script workers
func Run() {
	RunQuotaRemove()
}

// Stop - stops all script workers
func Stop() {
	log.Println(`Scripts stopping`)
	close(stopScripts)
}
