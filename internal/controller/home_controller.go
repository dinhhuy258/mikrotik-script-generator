package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController interface {
	Index(context *gin.Context)
}

type homeController struct{}

func NewHomeController() HomeController {
	return &homeController{}
}

type Script struct {
	Name        string
	Description string
	Route       string
}

var scripts = []Script{
	{Name: "Configure PPPoE", Description: "Set up PPPoE client configuration", Route: "/configure-pppoe"},
	{Name: "Configure WireGuard", Description: "Set up WireGuard VPN configuration", Route: "/configure-wireguard"},
	{Name: "Configure ECMP", Description: "Set up Equal-Cost Multi-Path routing", Route: "/configure-ecmp"},
}

func (_self *homeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"scripts": scripts,
	})
}
