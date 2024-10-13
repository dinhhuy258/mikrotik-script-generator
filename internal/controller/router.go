package controller

import (
	"mikrotik-script-generator/pkg/httpserver"
)

func SetRoutes(
	server httpserver.Interface,
	homeController HomeController,
	wireguardScriptController WireguardScriptController,
) {
	router := server.GetRouter()

	router.GET("/", homeController.Index)
	router.GET("/wireguard", wireguardScriptController.Index)
	router.POST("/wireguard", wireguardScriptController.GenerateMikrotikScript)
}
