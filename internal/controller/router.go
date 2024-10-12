package controller

import (
	"mikrotik-script-generator/pkg/httpserver"
)

func SetRoutes(
	server httpserver.Interface,
	homeController HomeController,
	scriptController ScriptController,
) {
	router := server.GetRouter()

	router.GET("/", homeController.Index)
	router.GET("/configure-wireguard", scriptController.ConfigWireGuard)
}
