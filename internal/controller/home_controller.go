package controller

import (
	"mikrotik-script-generator/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController interface {
	Index(context *gin.Context)
}

type homeController struct {
	homeService service.HomeService
}

func NewHomeController(homeService service.HomeService) HomeController {
	return &homeController{
		homeService: homeService,
	}
}

func (_self *homeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"IsHomePage":      true,
		"Title":           "Mikrotik Script Generator",
		"MikrotikScripts": _self.homeService.GetMikrotikScripts(),
	})
}
