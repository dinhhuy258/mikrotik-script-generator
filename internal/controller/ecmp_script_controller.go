package controller

import (
	"mikrotik-script-generator/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ecmpTitle            = "Configure ECMP"
	defaultECMPUsername  = "username"
	defaultECMPPassword  = "password"
	defaultECMPSessions  = 1
	defaultECMPInterface = "ether1"
)

type ECMPScriptController interface {
	BaseScriptController
}

type emcpScriptController struct{}

func NewECMPScriptController() ECMPScriptController {
	return &emcpScriptController{}
}

func (_self *emcpScriptController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "ecmp.html", gin.H{
		"Title": ecmpTitle,
		"FormData": model.ECMPFormData{
			Username:  defaultECMPUsername,
			Password:  defaultECMPPassword,
			Sessions:  defaultECMPSessions,
			Interface: defaultECMPInterface,
		},
	})
}

func (_self *emcpScriptController) GenerateMikrotikScript(c *gin.Context) {
	c.HTML(http.StatusOK, "ecmp.html", gin.H{
		"MikrotikScript": "hello",
	})
}
