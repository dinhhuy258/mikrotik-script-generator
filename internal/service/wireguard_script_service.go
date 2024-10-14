package service

import (
	"bytes"
	"html/template"
	"io"
	"mime/multipart"
	"net"

	"gopkg.in/ini.v1"
)

const defaultMTU = "1420"

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
	}
}

type WireguardScriptService interface {
	ParseConfig(cfgFile *multipart.FileHeader) (*WireguardConfig, error)
	GenerateScript(name string, listenPort int, configType string, wireguardCfg *WireguardConfig) (string, error)
}

type wireguardScriptService struct{}

func NewWireguardScriptService() WireguardScriptService {
	return &wireguardScriptService{}
}

func (_self *wireguardScriptService) GenerateScript(
	name string,
	listenPort int,
	configType string,
	wireguardCfg *WireguardConfig,
) (string, error) {
	tmpl, err := template.ParseFiles("internal/service/mikrotik/wireguard_script.tmpl")
	if err != nil {
		return "", err
	}

	var script bytes.Buffer

	err = tmpl.Execute(&script, map[string]interface{}{
		"Name":       name,
		"ListenPort": listenPort,
		"ConfigType": configType,
		"Wireguard":  wireguardCfg,
	})
	if err != nil {
		return "", err
	}

	return script.String(), nil
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
	address := interfaceSection.Key("Address").String()

	mtu := interfaceSection.Key("MTU").String()
	if mtu == "" {
		mtu = defaultMTU
	}

	peerSection := cfg.Section("Peer")
	publicKey := peerSection.Key("PublicKey").String()
	endpoint := peerSection.Key("Endpoint").String()
	presharedKey := peerSection.Key("PresharedKey").String()

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
		}{
			PublicKey:    publicKey,
			EndpointHost: endpointHost,
			EndpointPort: endpointPort,
			PresharedKey: presharedKey,
		},
	}, nil
}
