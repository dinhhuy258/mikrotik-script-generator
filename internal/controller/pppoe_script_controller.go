package controller

import (
	"mikrotik-script-generator/internal/model"
	"mikrotik-script-generator/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	pppoeTitle             = "Configure PPPoE"
	defaultPPPoEUsername   = "username"
	defaultPPPoEPassword   = "password"
	defaultPPPoEInterface  = "ether1"
	defaultPPPoELANNetwork = "192.168.0.1/24"
)

type PPPoEScriptController interface {
	BaseScriptController
}

type pppoeScriptController struct {
	pppoeScriptService service.PPPoEScriptService
}

func NewPPPoEScriptController(pppoeScriptService service.PPPoEScriptService) PPPoEScriptController {
	return &pppoeScriptController{
		pppoeScriptService: pppoeScriptService,
	}
}

func (_self *pppoeScriptController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "pppoe.html", gin.H{
		"Title": pppoeTitle,
		"FormData": model.PPPoEFormData{
			Username:   defaultPPPoEUsername,
			Password:   defaultPPPoEPassword,
			Interface:  defaultPPPoEInterface,
			LANNetwork: defaultPPPoELANNetwork,
		},
	})
}

func (_self *pppoeScriptController) GenerateMikrotikScript(c *gin.Context) {
	var pppoeFormData model.PPPoEFormData
	if err := c.ShouldBind(&pppoeFormData); err != nil {
		c.HTML(http.StatusOK, "pppoe.html", gin.H{
			"Error": "There was an error processing your request",
		})

		return
	}

	mikrotikScript, err := _self.pppoeScriptService.GenerateScript(
		pppoeFormData.Username,
		pppoeFormData.Password,
		pppoeFormData.Interface,
		pppoeFormData.LANNetwork,
	)
	if err != nil {
		c.HTML(http.StatusOK, "pppoe.html", gin.H{
			"Error": "There was an error generating your Mikrotik script",
		})

		return
	}

	c.HTML(http.StatusOK, "pppoe.html", gin.H{
		"MikrotikScript": mikrotikScript,
	})
}
