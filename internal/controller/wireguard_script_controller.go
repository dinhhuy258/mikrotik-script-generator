package controller

import (
	"mikrotik-script-generator/internal/model"
	"mikrotik-script-generator/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	title             = "Configure WireGuard"
	defaultListenPort = 13231
	defaultName       = "mikrotik"
)

type WireguardScriptController interface {
	Index(c *gin.Context)
	GenerateMikrotikScript(c *gin.Context)
}

type wireguardScriptController struct {
	wireguardScriptService service.WireguardScriptService
}

func NewWireguardScriptController(wireguardScriptService service.WireguardScriptService) WireguardScriptController {
	return &wireguardScriptController{
		wireguardScriptService: wireguardScriptService,
	}
}

func (_self *wireguardScriptController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "wireguard.html", gin.H{
		"Title": title,
		"FormData": model.WireGuardFormData{
			Name:       defaultName,
			ListenPort: defaultListenPort,
		},
	})
}

func (_self *wireguardScriptController) GenerateMikrotikScript(c *gin.Context) {
	var wireGuardFormData model.WireGuardFormData
	if err := c.ShouldBind(&wireGuardFormData); err != nil {
		c.HTML(http.StatusOK, "wireguard.html", gin.H{
			"Title": title,
			"FormData": model.WireGuardFormData{
				Name:       defaultName,
				ListenPort: defaultListenPort,
			},
			"Error": "There was an error processing your request",
		})

		return
	}

	wireGuardConfig, err := _self.wireguardScriptService.ParseConfig(wireGuardFormData.ConfigFile)
	if err != nil {
		c.HTML(http.StatusOK, "wireguard.html", gin.H{
			"Title": title,
			"FormData": model.WireGuardFormData{
				Name:       defaultName,
				ListenPort: defaultListenPort,
			},
			"Error": "There was an error parsing your WireGuard configuration file",
		})

		return
	}

	mikrotikScript, err := _self.wireguardScriptService.GenerateScript(
		wireGuardFormData.Name,
		wireGuardFormData.ListenPort,
		wireGuardFormData.ConfigType,
		wireGuardConfig,
	)
	if err != nil {
		c.HTML(http.StatusOK, "wireguard.html", gin.H{
			"Title": title,
			"FormData": model.WireGuardFormData{
				Name:       defaultName,
				ListenPort: defaultListenPort,
			},
			"Error": "There was an error generating your Mikrotik script",
		})

		return
	}

	c.HTML(http.StatusOK, "wireguard.html", gin.H{
		"MikrotikScript": mikrotikScript,
	})
}
