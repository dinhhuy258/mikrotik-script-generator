package controller

import (
	"mikrotik-script-generator/internal/model"
	"mikrotik-script-generator/internal/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

const (
	pppoeTitle                 = "Configure PPPoE"
	defaultPPPoEUsername       = "username"
	defaultPPPoEPassword       = "password"
	defaultPPPoEInterface      = "ether1"
	defaultPPPoEBridgeLANPorts = "ether2,ether3,ether4"
	defaultPPPoEGateway        = "192.168.0.1"
	defaultPPPoELANNetwork     = "192.168.0.0/24"
	defaultPPPoEDHCPRange      = "192.168.0.10-192.168.0.254"
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
			Username:      defaultPPPoEUsername,
			Password:      defaultPPPoEPassword,
			Interface:     defaultPPPoEInterface,
			BridgeLANPort: defaultPPPoEBridgeLANPorts,
			Gateway:       defaultPPPoEGateway,
			LANNetwork:    defaultPPPoELANNetwork,
			DHCPRange:     defaultPPPoEDHCPRange,
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

	pppoeFormData.Interface = strings.TrimSpace(pppoeFormData.Interface)
	bridgeLANPorts := strings.Split(pppoeFormData.BridgeLANPort, ",")
	bridgeLANPorts = lo.Map(bridgeLANPorts, func(port string, _ int) string {
		return strings.TrimSpace(port)
	})

	if _, invalidBridgeLanPort := lo.Find(bridgeLANPorts, func(port string) bool {
		return port == pppoeFormData.Interface
	}); invalidBridgeLanPort {
		c.HTML(http.StatusOK, "pppoe.html", gin.H{
			"Error": "The bridge LAN port must not be the same as the PPPoE interface",
		})

		return
	}

	mikrotikScript, err := _self.pppoeScriptService.GenerateScript(
		pppoeFormData.Username,
		pppoeFormData.Password,
		pppoeFormData.Interface,
		bridgeLANPorts,
		pppoeFormData.Gateway,
		pppoeFormData.LANNetwork,
		pppoeFormData.DHCPRange,
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
