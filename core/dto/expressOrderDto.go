package dto

import "encoding/xml"

type OrderExpressDto struct {
	BizData     string `json:"bizData"`
	ServiceType string `json:"serviceType"`
	PartnerID   string `json:"partnerID"`
	Sign        string `json:"sign"`
}

type BaiShiRequestDto struct {
	XMLName       xml.Name `xml:"request"`
	CustomerCode  string   `xml:"customerCode"`
	CustomerName  string   `xml:"customerName"`
	ProjectCode   string   `xml:"projectCode"`
	TransportMode string   `xml:"transportMode"`
	VehicleModel  string   `xml:"vehicleModel"`

	LogisticsCode    string `xml:"logisticsCode"`
	Remark           string `xml:"remark"`
	OrderDescription string `xml:"orderDescription"`
	ShippingOrderNo  string `xml:"shippingOrderNo"`
	OrderCode        string `xml:"orderCode"`

	GoodsValue  string `xml:"goodsValue"`
	CheapAmount string `xml:"cheapAmount"`
	NumberType  string `xml:"type"`
	Value       string `xml:"value"`

	Sender   BaiShiSenderDto     `xml:"sender"`
	Receiver BaiShiReceiverDto   `xml:"receiver"`
	Items    []*BaiShiItemDtoNew `xml:items`
}

type BaiShiSenderDto struct {
	XMLName  xml.Name `xml:"sender"`
	Name     string   `xml:"name"`
	Province string   `xml:"province"`
	City     string   `xml:"city"`
	District string   `xml:"district"`
	Address  string   `xml:"address"`

	ContactName     string `xml:"contactName"`
	Phone           string `xml:"phone"`
	EarlyPickupTime string `xml:"earlyPickupTime"`
	LatePickupTime  string `xml:"latePickupTime"`
}

type BaiShiReceiverDto struct {
	XMLName  xml.Name `xml:"receiver"`
	Name     string   `xml:"name"`
	Province string   `xml:"province"`
	City     string   `xml:"city"`
	District string   `xml:"district"`
	Address  string   `xml:"address"`

	ContactName       string `xml:"contactName"`
	Phone             string `xml:"phone"`
	EarlyDeliveryTime string `xml:"earlyDeliveryTime"`
	LateDeliveryTime  string `xml:"lateDeliveryTime"`
}

type BaiShiItemDtoNew struct {
	Item BaiShiItemDto `xml:"item"`
}

type BaiShiItemDto struct {
	LineNo       string `lineNo`
	ItemCode     string `itemCode`
	ItemName     string `itemName`
	ItemCount    string `itemCount`
	PackageCount string `packageCount`

	PackageUomCode string `packageUomCode`
	Weight         string `weight`
	Volume         string `volume`
	UnitPrice      string `unitPrice`
	DeclaredValue  string `declaredValue`

	VolumeWeight string `volumeWeight`
	Remark       string `remark`
}

type BaiShiAdditionalDto struct {
	addType string ``
}
