package controller

import (
	"mikrotik-script-generator/pkg/httpserver"
)

func SetRoutes(
	server httpserver.Interface,
	homeController HomeController,
) {
	router := server.GetRouter()

	router.GET("/", homeController.Index)
}
