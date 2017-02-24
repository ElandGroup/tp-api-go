package config

import "github.com/jinzhu/configor"

var Config = struct {
	Sample struct {
		Conn string `json:"Conn"`
	} `json:"Sample"`

	BaiShi struct {
		OrderExpressUrl string `json:"OrderExpressUrl"`
		PartnerId       string `json:"PartnerId"`
		PartnerKey      string `json:"PartnerKey"`
		CustomerCode    string `json:"CustomerCode"`
		ProjectCode     string `json:"ProjectCode"`
		ServiceType     string `json:"ServiceType"`
		TransportMode   string `json:"TransportMode"`
	}
}{}

func InitConfig(cfg string) {
	// configor.Load(&Config, "config/config.json")
	configor.Load(&Config, cfg)
}
