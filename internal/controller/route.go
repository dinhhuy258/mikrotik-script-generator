package controller

import (
	"mikrotik-script-generator/pkg/httpserver"
)

func SetRoutes(
	server httpserver.Interface,
	homeController HomeController,
	wireguardScriptController WireguardScriptController,
	ecmpScriptController ECMPScriptController,
	pppoeScriptController PPPoEScriptController,
) {
	router := server.GetRouter()

	router.GET("/", homeController.Index)
	router.GET("/wireguard", wireguardScriptController.Index)
	router.POST("/wireguard", wireguardScriptController.GenerateMikrotikScript)
	router.GET("/ecmp", ecmpScriptController.Index)
	router.POST("/ecmp", ecmpScriptController.GenerateMikrotikScript)
	router.GET("/pppoe", pppoeScriptController.Index)
	router.POST("/pppoe", pppoeScriptController.GenerateMikrotikScript)
}
