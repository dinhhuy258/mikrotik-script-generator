package model

type IPRoutingFormData struct {
	IPAddresses  []string          `form:"ipAddresses"`
	Gateway      string            `form:"gateway"`
	RoutingTable string            `form:"routingTable"`
	Errors       map[string]string `form:"errors"`
}
