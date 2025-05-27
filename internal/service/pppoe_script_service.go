package service

type PPPoEScriptService interface {
	GenerateScript(
		username,
		password,
		interfaceName string,
		bridgeLANPorts []string,
		gateway,
		lanNetwork,
		dhcpRange string,
	) (string, error)
}

type pppoeScriptService struct {
	BaseScriptGenerator
}

func NewPPPoEScriptService() PPPoEScriptService {
	return &pppoeScriptService{}
}

func (_self *pppoeScriptService) GenerateScript(
	username,
	password,
	interfaceName string,
	bridgeLANPorts []string,
	gateway,
	lanNetwork,
	dhcpRange string,
) (string, error) {
	return _self.GenerateScriptFromTemplate("internal/service/mikrotik/pppoe_script.tmpl", map[string]any{
		"Username":       username,
		"Password":       password,
		"Interface":      interfaceName,
		"BridgeLANPorts": bridgeLANPorts,
		"Gateway":        gateway,
		"LANNetwork":     lanNetwork,
		"DHCPRange":      dhcpRange,
	})
}
