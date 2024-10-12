package model

type WireGuardFormData struct {
	Name           string
	ListenPort     int
	MTU            int
	ConfigType     string
	ConfigFileData string
	Errors         map[string]string
}
