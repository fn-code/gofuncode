package main

import (
	"app/route"
	"app/shared/server"
)

func main() {

	server.Run(route.LoadHTTP())
}
