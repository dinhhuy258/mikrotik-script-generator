package controller

import (
	"mikrotik-script-generator/internal/model"
	"mikrotik-script-generator/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ecmpTitle             = "Configure ECMP"
	defaultECMPUsername   = "username"
	defaultECMPPassword   = "password"
	defaultECMPSessions   = 1
	defaultECMPInterface  = "ether1"
	defaultECMPLANNetwork = "192.168.0.1/24"
)

type ECMPScriptController interface {
	BaseScriptController
}

type emcpScriptController struct {
	ecmpScriptService service.ECMPScriptService
}

func NewECMPScriptController(ecmpScriptService service.ECMPScriptService) ECMPScriptController {
	return &emcpScriptController{
		ecmpScriptService: ecmpScriptService,
	}
}

func (_self *emcpScriptController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "ecmp.html", gin.H{
		"Title": ecmpTitle,
		"FormData": model.ECMPFormData{
			Username:   defaultECMPUsername,
			Password:   defaultECMPPassword,
			Sessions:   defaultECMPSessions,
			Interface:  defaultECMPInterface,
			LANNetwork: defaultECMPLANNetwork,
		},
	})
}

func (_self *emcpScriptController) GenerateMikrotikScript(c *gin.Context) {
	var ecmpFormData model.ECMPFormData
	if err := c.ShouldBind(&ecmpFormData); err != nil {
		c.HTML(http.StatusOK, "ecmp.html", gin.H{
			"Error": "There was an error processing your request",
		})

		return
	}

	sessions := make([]int, ecmpFormData.Sessions)
	for i := 0; i < ecmpFormData.Sessions; i++ {
		sessions[i] = i + 1
	}

	mikrotikScript, err := _self.ecmpScriptService.GenerateScript(
		ecmpFormData.Username,
		ecmpFormData.Password,
		sessions,
		ecmpFormData.Interface,
		ecmpFormData.LANNetwork,
	)
	if err != nil {
		c.HTML(http.StatusOK, "ecmp.html", gin.H{
			"Error": "There was an error generating your Mikrotik script",
		})

		return
	}

	c.HTML(http.StatusOK, "ecmp.html", gin.H{
		"MikrotikScript": mikrotikScript,
	})
}
