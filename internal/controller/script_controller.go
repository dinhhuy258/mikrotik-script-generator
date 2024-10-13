package controller

import (
	"mikrotik-script-generator/internal/model"
	"mikrotik-script-generator/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScriptController interface {
	ConfigWireGuard(context *gin.Context)
	ConfigWireGuardPost(context *gin.Context)
}

type scriptController struct {
	scriptService service.ScriptService
}

func NewScriptController(scriptService service.ScriptService) ScriptController {
	return &scriptController{
		scriptService: scriptService,
	}
}

func (_self *scriptController) ConfigWireGuard(c *gin.Context) {
	c.HTML(http.StatusOK, "configure-wireguard.html", gin.H{
		"formData": model.WireGuardFormData{
			Name:       "mikrotik",
			MTU:        1420,
			ListenPort: 13231,
		},
	})
}

func (_self *scriptController) ConfigWireGuardPost(c *gin.Context) {
	c.HTML(http.StatusOK, "script.html", gin.H{
		"generatedScript": "Hello world",
	})
}
