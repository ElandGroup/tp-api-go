package dto

type BaiShiQueryDto struct {
	orderCode      string `json:"orderCode"`
	customerCode   string `json:"customerCode"`
	createTimeFrom string `json:"createTimeFrom"`
	createTimeTo   string `json:"createTimeTo"`
	shipmentCode   string `json:"shipmentCode"`
}
