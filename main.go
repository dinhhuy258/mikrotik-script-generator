package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Script struct {
	Name        string
	Description string
	Route       string
}

type FormData struct {
	Username string
	Password string
	Interface string
	Errors map[string]string
}

var scripts = []Script{
	{Name: "Configure PPPoE", Description: "Set up PPPoE client configuration", Route: "/configure-pppoe"},
	{Name: "Configure WireGuard", Description: "Set up WireGuard VPN configuration", Route: "/configure-wireguard"},
	{Name: "Configure ECMP", Description: "Set up Equal-Cost Multi-Path routing", Route: "/configure-ecmp"},
}

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", homePage)
	r.GET("/configure-pppoe", configurePPPoEPage)
	r.POST("/configure-pppoe", handlePPPoESubmission)

	r.Run(":8080")
}

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"scripts": scripts,
	})
}

func configurePPPoEPage(c *gin.Context) {
	c.HTML(http.StatusOK, "configure-pppoe.html", gin.H{
		"formData": FormData{},
	})
}

func handlePPPoESubmission(c *gin.Context) {
	username := strings.TrimSpace(c.PostForm("username"))
	password := strings.TrimSpace(c.PostForm("password"))
	interfaceName := strings.TrimSpace(c.PostForm("interface"))

	formData := FormData{
		Username: username,
		Password: password,
		Interface: interfaceName,
		Errors: make(map[string]string),
	}

	if username == "" {
		formData.Errors["username"] = "Username is required"
	}
	if password == "" {
		formData.Errors["password"] = "Password is required"
	}
	if interfaceName == "" {
		formData.Errors["interface"] = "Interface is required"
	}

	if len(formData.Errors) > 0 {
		c.HTML(http.StatusBadRequest, "configure-pppoe.html", gin.H{
			"formData": formData,
		})
		return
	}

	script := generatePPPoEScriptContent(username, password, interfaceName)

	c.HTML(http.StatusOK, "configure-pppoe.html", gin.H{
		"formData": formData,
		"generatedScript": script,
	})
}

func generatePPPoEScriptContent(username, password, interfaceName string) string {
	return `/interface pppoe-client
add name=pppoe-out1 interface=` + interfaceName + ` user=` + username + ` password=` + password + ` disabled=no

/ip route
add distance=1 gateway=pppoe-out1

/ip dns
set servers=8.8.8.8,8.8.4.4

/ip firewall nat
add chain=srcnat out-interface=pppoe-out1 action=masquerade`
}
