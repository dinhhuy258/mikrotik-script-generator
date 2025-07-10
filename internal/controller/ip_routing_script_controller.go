package controller

import (
	"mikrotik-script-generator/internal/model"
	"mikrotik-script-generator/internal/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ipRoutingTitle = "IP Routing Configuration"
)

type IPRoutingScriptController interface {
	BaseScriptController
}

type ipRoutingScriptController struct {
	ipRoutingScriptService service.IPRoutingScriptService
}

func NewIPRoutingScriptController(ipRoutingScriptService service.IPRoutingScriptService) IPRoutingScriptController {
	return &ipRoutingScriptController{
		ipRoutingScriptService: ipRoutingScriptService,
	}
}

func (_self *ipRoutingScriptController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "ip_routing.html", gin.H{
		"Title": ipRoutingTitle,
		"FormData": model.IPRoutingFormData{
			Gateway:      "cloudflare-wireguard",
			RoutingTable: "telegram",
			IPAddresses: []string{
				"91.108.4.0/22",
				"91.108.8.0/22",
				"91.108.12.0/22",
				"91.108.16.0/22",
				"91.108.20.0/22",
				"91.108.56.0/22",
				"95.161.64.0/20",
				"149.154.160.0/20",
				"149.154.164.0/22",
				"149.154.168.0/22",
				"149.154.172.0/22",
			},
		},
	})
}

func (_self *ipRoutingScriptController) GenerateMikrotikScript(c *gin.Context) {
	var ipRoutingFormData model.IPRoutingFormData
	if err := c.ShouldBind(&ipRoutingFormData); err != nil {
		c.HTML(http.StatusOK, "ip_routing.html", gin.H{
			"Error": "There was an error processing your request",
		})

		return
	}

	// Parse IP addresses from textarea input
	ipAddressesText := c.PostForm("ipAddressesText")
	if ipAddressesText == "" {
		c.HTML(http.StatusOK, "ip_routing.html", gin.H{
			"Error": "Please provide at least one IP address",
		})

		return
	}

	var cleanedIPs []string

	// Split by newlines and clean up
	ipAddresses := strings.Split(ipAddressesText, "\n")
	for _, ip := range ipAddresses {
		ip = strings.TrimSpace(ip)
		if ip != "" {
			cleanedIPs = append(cleanedIPs, ip)
		}
	}

	if len(cleanedIPs) == 0 {
		c.HTML(http.StatusOK, "ip_routing.html", gin.H{
			"Error": "Please provide at least one valid IP address",
		})

		return
	}

	ipRoutingFormData.IPAddresses = cleanedIPs

	mikrotikScript, err := _self.ipRoutingScriptService.GenerateScript(ipRoutingFormData)
	if err != nil {
		c.HTML(http.StatusOK, "ip_routing.html", gin.H{
			"Error": "There was an error generating your Mikrotik script",
		})

		return
	}

	reverseScript, err := _self.ipRoutingScriptService.GenerateReverseScript(ipRoutingFormData)
	if err != nil {
		c.HTML(http.StatusOK, "ip_routing.html", gin.H{
			"Error": "There was an error generating your reverse script",
		})
		return
	}

	c.HTML(http.StatusOK, "ip_routing.html", gin.H{
		"MikrotikScript": mikrotikScript,
		"ReverseScript":  reverseScript,
		"FormData":       ipRoutingFormData,
	})
}
