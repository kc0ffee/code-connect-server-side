package main

import (
	"github.com/kc0ffee/server/server"
)

func main() {
	e := server.NewAPIServer()
	server.StartServer(e, 10000)
}
