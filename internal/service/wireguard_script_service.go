package service

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"strings"

	"github.com/samber/lo"
	"gopkg.in/ini.v1"
)

const (
	defaultMTU       = "1420"
	defaultAllowedIP = "0.0.0.0/0"
)

type WireguardConfig struct {
	Interface struct {
		PrivateKey string
		Address    string
		MTU        string
	}
	Peer struct {
		PublicKey    string
		EndpointHost string
		EndpointPort string
		PresharedKey string
		AllowedIPs   []string
	}
}

type WireguardScriptService interface {
	ParseConfig(cfgFile *multipart.FileHeader) (*WireguardConfig, error)
	GenerateScript(name string, listenPort int, configType string, wireguardCfg *WireguardConfig) (string, error)
}

type wireguardScriptService struct {
	BaseScriptGenerator
}

func NewWireguardScriptService() WireguardScriptService {
	return &wireguardScriptService{}
}

func (_self *wireguardScriptService) GenerateScript(
	name string,
	listenPort int,
	configType string,
	wireguardCfg *WireguardConfig,
) (string, error) {
	return _self.GenerateScriptFromTemplate("internal/service/mikrotik/wireguard_script.tmpl",
		map[string]interface{}{
			"Name":       name,
			"ListenPort": listenPort,
			"ConfigType": configType,
			"Wireguard":  wireguardCfg,
		})
}

func (_self *wireguardScriptService) ParseConfig(cfgFile *multipart.FileHeader) (*WireguardConfig, error) {
	file, err := cfgFile.Open()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	cfgBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	cfgReader := bytes.NewReader(cfgBytes)

	cfg, err := ini.Load(cfgReader)
	if err != nil {
		return nil, err
	}

	interfaceSection := cfg.Section("Interface")
	privateKey := interfaceSection.Key("PrivateKey").String()
	// Split the string by comma and remove any leading/trailing whitespaces
	addresses := strings.Split(interfaceSection.Key("Address").String(), ",")
	addresses = lo.Map(addresses, func(addr string, _ int) string {
		return strings.TrimSpace(addr)
	})
	// Only allow IPv4 addresses
	addresses = lo.Filter(addresses, func(addr string, _ int) bool {
		return isIPv4(addr)
	})
	if len(addresses) == 0 {
		return nil, fmt.Errorf("No valid IPv4 address found in the Address field")
	}
	// Pick the first address
	address := addresses[0]

	mtu := interfaceSection.Key("MTU").String()
	if mtu == "" {
		mtu = defaultMTU
	}

	peerSection := cfg.Section("Peer")
	publicKey := peerSection.Key("PublicKey").String()
	endpoint := peerSection.Key("Endpoint").String()
	presharedKey := peerSection.Key("PresharedKey").String()
	allowedIPsStr := peerSection.Key("AllowedIPs").String()
	// Split the string by comma and remove any leading/trailing whitespaces
	allowedIPs := strings.Split(allowedIPsStr, ",")
	// Remove any leading/trailing whitespaces
	allowedIPs = lo.Map(allowedIPs, func(ip string, _ int) string {
		return strings.TrimSpace(ip)
	})
	// Only allow IPv4 addresses
	allowedIPs = lo.Filter(allowedIPs, func(ip string, _ int) bool {
		return isIPv4(ip)
	})
	if len(allowedIPs) == 0 {
		allowedIPs = []string{defaultAllowedIP}
	}

	endpointHost, endpointPort, err := net.SplitHostPort(endpoint)
	if err != nil {
		return nil, err
	}

	return &WireguardConfig{
		Interface: struct {
			PrivateKey string
			Address    string
			MTU        string
		}{
			PrivateKey: privateKey,
			Address:    address,
			MTU:        mtu,
		},
		Peer: struct {
			PublicKey    string
			EndpointHost string
			EndpointPort string
			PresharedKey string
			AllowedIPs   []string
		}{
			PublicKey:    publicKey,
			EndpointHost: endpointHost,
			EndpointPort: endpointPort,
			PresharedKey: presharedKey,
			AllowedIPs:   allowedIPs,
		},
	}, nil
}

func isIPv4(address string) bool {
	// Check if address has a CIDR notation (subnet)
	if strings.Contains(address, "/") {
		ip, _, err := net.ParseCIDR(address)

		return err == nil && ip.To4() != nil
	}

	ip := net.ParseIP(address)

	return ip != nil && ip.To4() != nil
}
