package main

import (
	"fmt"
	"mystly/internal/core"
	"mystly/internal/routes"
)

func main() {
	// TODO: Config
	server := core.NewServer()
	routes.RegisterHandlers(server)

	server.Run(fmt.Sprintf(":%d", server.Config.Port), server.Config.PostgresDSN)
}
