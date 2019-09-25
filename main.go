package main

import (
	"alif/quota/api"
	"alif/quota/api/scripts"
)

func main() {
	go scripts.Run()
	defer scripts.Stop()

	api.Run()
}
