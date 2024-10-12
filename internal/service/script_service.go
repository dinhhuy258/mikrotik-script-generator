package service

type ScriptService interface {
	ConfigureWireGuard() string
}

type scriptService struct{}

func NewScriptService() ScriptService {
	return &scriptService{}
}

func (_self *scriptService) ConfigureWireGuard() string {
	return "Configure WireGuard"
}
