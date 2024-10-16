package controller

import (
	"github.com/gin-gonic/gin"
)

type BaseScriptController interface {
	Index(c *gin.Context)
	GenerateMikrotikScript(c *gin.Context)
}
