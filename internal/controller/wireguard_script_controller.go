package controller

import (
	"mikrotik-script-generator/internal/model"
	"mikrotik-script-generator/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WireguardScriptController interface {
	Index(c *gin.Context)
	GenerateMikrotikScript(c *gin.Context)
}

type wireguardScriptController struct {
	scriptService service.ScriptService
}

func NewWireguardScriptController(scriptService service.ScriptService) WireguardScriptController {
	return &wireguardScriptController{
		scriptService: scriptService,
	}
}

func (_self *wireguardScriptController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "wireguard.html", gin.H{
		"formData": model.WireGuardFormData{
			Name:       "mikrotik",
			MTU:        1420,
			ListenPort: 13231,
		},
	})
}

func (_self *wireguardScriptController) GenerateMikrotikScript(c *gin.Context) {
	c.HTML(http.StatusOK, "mikrotik_script.html", gin.H{
		"mikrotikScript": "Hello world",
	})
}
